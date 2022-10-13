package routers

import (
	"glofox-task/controllers"

	"github.com/go-chi/chi/v5"
)

func AddClassRouter(router chi.Router, classController *controllers.ClassController) {
	router.Post("/api/class", classController.Create)
}
