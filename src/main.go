package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"glofox-task/controllers"
	"glofox-task/repositories"
	"glofox-task/routers"
)

func main() {

	// Initialize repositories and controllers
	repos := repositories.InitRepositories("db")
	ctrls := controllers.InitControllers(repos)

	// Creates routing
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Adds Routers
	routers.AddBookingRouter(router, ctrls.BookingController)
	routers.AddClassRouter(router, ctrls.ClassController)

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
