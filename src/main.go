package main

import (
	"log"
	"net/http"
	"os"

	_ "glofox-task/docs" // Swagger requires this

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"glofox-task/controllers"
	"glofox-task/database"
	"glofox-task/repositories"
	"glofox-task/routers"
)

// @title Glofox-task App Swagger
// @version 1.0
// @description This is a task for an interview

// @contact.name Francisco Barao Santos
// @contact.email s.franciscobarao@gmail.com

// @BasePath /api/
func main() {

	// Connect to Database
	db, err := database.Connect()
	if err != nil {
		log.Println("Error occurred while connecting to database")
		return
	}

	// Initialize repositories and controllers
	repos := repositories.InitRepositories(db)
	ctrls := controllers.InitControllers(repos)

	// Creates routing
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Adds Routers
	routers.AddBookingRouter(router, ctrls.BookingController)
	routers.AddClassRouter(router, ctrls.ClassController)

	// Documentation for developers
	router.Get("/swagger/*", httpSwagger.Handler())

	// Starts server
	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		log.Println("Error occurred while fetching Port")
		return
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println("Error occured while creating Server" + err.Error())
		return
	}
	log.Println("Server is Running on localhost:" + port)
}
