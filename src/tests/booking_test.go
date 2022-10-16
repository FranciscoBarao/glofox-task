package tests

import (
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/steinfletcher/apitest"

	"glofox-task/controllers"
	"glofox-task/database"
	"glofox-task/repositories"
	"glofox-task/routers"
)

var router *chi.Mux

func init() {
	log.Println("Setup Starting")

	// Connect to Database
	db, err := database.Connect()
	if err != nil {
		log.Println("Error occurred while connecting to database")
		return
	}

	// Set Repositories & Controllers
	repos := repositories.InitRepositories(db)
	ctrls := controllers.InitControllers(repos)

	// Creates Router
	router = chi.NewRouter()
	routers.AddBookingRouter(router, ctrls.BookingController)

	log.Println("Setup Complete")
}

/*Success Tests*/

func TestCreateBookingSuccess(t *testing.T) {

	// To succeed a class in that date must exist. Assuming we run class tests first, it exists.
	// There should be some way of mocking a class existence.

	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"name":"JohnDoe", "date":"2022-01-01"}`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetAllBookingSuccess(t *testing.T) {
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Get("/api/booking").
		Expect(t).
		Status(http.StatusOK).
		End()
}

/*Failure Tests*/
func TestCreateBookingFailureNoClass(t *testing.T) {
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"name":"JohnDoe", "date":"2022-02-01"}`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateBookingFailureOverbooking(t *testing.T) {
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"name":"JohnDoe", "date":"2022-01-01"}`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusConflict).
		End()
}

func TestCreateBookingJsonFailures(t *testing.T) {
	// Several Json Objects on the body
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`[{"name":"JohnDoe", "date":"2022-01-01"}, {"name":"JohnDoe", "date":"2022-01-01"}]`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Malformed Json
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"name :"JohnDoe", "date":"2022-01-01"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Unmarshall type error
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"name":1, "date":"2022-01-01"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Unknown Field
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(`{"test":"test", "name":"JohnDoe", "date":"2022-01-01"}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Empty Body
	apitest.New().
		HandlerFunc(router.ServeHTTP).
		Post("/api/booking").
		JSON(``).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestCreateBookingValidStructFailures(t *testing.T) {
	//  <<<< field - Name >>>>
	apitest.New(). // Invalid Struct -> NOT maxstringlength(50)
			HandlerFunc(router.ServeHTTP).
			Post("/api/booking").
			JSON(`{"name":"JohnDoessssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "date":"2012-04-23"}`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

	apitest.New(). // Invalid Struct -> NOT alphanum
			HandlerFunc(router.ServeHTTP).
			Post("/api/booking").
			JSON(`{"name":"John ?!@#_123 Doe", "date":"2012-04-23"}`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

	//  <<<< field - Date >>>>
	apitest.New(). // Invalid Struct -> NOT ("yyyy-mm-dd")
			HandlerFunc(router.ServeHTTP).
			Post("/api/booking").
			JSON(`{"name":"JohnDoe", "date":"2006-01-02T15:04:05"}`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
}
