package controllers

import (
	"net/http"

	"github.com/unrolled/render"

	"glofox-task/middleware"
	"glofox-task/models"
	"glofox-task/repositories"
	"glofox-task/utils"
)

// Declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type classRepository interface {
	Create(class *models.Class) error
	ReadAll() ([]models.Class, error)
}

type ClassController struct {
	repo classRepository
}

func InitClassController(bookingRepo *repositories.ClassRepository) *ClassController {
	return &ClassController{
		repo: bookingRepo,
	}
}

func (controller *ClassController) Create(w http.ResponseWriter, r *http.Request) {

	// Body to Class Struct
	var class models.Class
	if err := utils.DecodeJSONBody(w, r, &class); err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	// Validate Class input
	if err := utils.ValidateStruct(&class); err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	// Calls create on repository
	if err := controller.repo.Create(&class); err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	render.New().JSON(w, http.StatusOK, class)
}

func (controller *ClassController) GetAll(w http.ResponseWriter, r *http.Request) {

	// Calls create on repository
	classes, err := controller.repo.ReadAll()
	if err != nil {
		middleware.ErrorHandler(w, err)
		return
	}

	render.New().JSON(w, http.StatusOK, classes)
}
