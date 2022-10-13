package controllers

import (
	"net/http"

	"github.com/unrolled/render"

	"glofox-task/repositories"
)

// Declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type classRepository interface {
	Create() (string, error)
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

	test, _ := controller.repo.Create()

	render.New().JSON(w, http.StatusOK, test)
}
