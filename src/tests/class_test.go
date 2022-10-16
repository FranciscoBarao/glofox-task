package tests

import (
	"glofox-task/controllers"
	"glofox-task/database"
	"glofox-task/repositories"
	"glofox-task/routers"
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/steinfletcher/apitest"
)

var route *chi.Mux

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
	route = chi.NewRouter()
	routers.AddClassRouter(route, ctrls.ClassController)

	log.Println("Setup Complete")
}

/*Success Tests*/

func TestCreateClassSuccess(t *testing.T) {
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":1 }`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetAllClassSuccess(t *testing.T) {
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Get("/api/class").
		Expect(t).
		Status(http.StatusOK).
		End()
}

/*Failure Tests*/
func TestCreateClassFailureOverlap(t *testing.T) {
	// Overlap last day
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"name":"Aerobics", "start_date":"2022-01-03", "end_date":"2022-01-05", "capacity":2 }`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusConflict).
		End()

	// Overlap everything
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusConflict).
		End()
}

func TestCreateClassJsonFailures(t *testing.T) {
	// Several Json Objects on the body
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`[{"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }, {"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }]`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Malformed Json
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"name:"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Unmarshall type error
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"name": 1, "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Unknown Field
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(`{"test":"test", "name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Empty Body
	apitest.New().
		HandlerFunc(route.ServeHTTP).
		Post("/api/class").
		JSON(``).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestCreateClassValidStructFailures(t *testing.T) {
	//  <<<< field - Name >>>>
	apitest.New(). // Invalid Struct -> NOT maxstringlength(50)
			HandlerFunc(route.ServeHTTP).
			Post("/api/class").
			JSON(`{"name":"Aerobicsaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

	apitest.New(). // Invalid Struct -> NOT alphanum
			HandlerFunc(route.ServeHTTP).
			Post("/api/class").
			JSON(`{"name":"Aerobic?!@#_! s", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

	//  <<<< field - StartDate >>>>
	apitest.New(). // Invalid Struct -> NOT ("yyyy-mm-dd")
			HandlerFunc(route.ServeHTTP).
			Post("/api/class").
			JSON(`{"name":"Aerobics", "start_date":"2006-01-02T15:04:05", "end_date":"2022-01-03", "capacity":2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

	//  <<<< field - EndDate >>>>
	apitest.New(). // Invalid Struct -> NOT ("yyyy-mm-dd")
			HandlerFunc(route.ServeHTTP).
			Post("/api/class").
			JSON(`{"name":"Aerobics", "start_date":"2022-01-03", "end_date":"2006-01-02T15:04:05", "capacity":2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

		//  <<<< field - Capacity >>>>
	apitest.New().
		HandlerFunc(route.ServeHTTP). // Invalid Struct -> NOT 1|100
		Post("/api/class").
		JSON(`{"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":200 }`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	apitest.New().
		HandlerFunc(route.ServeHTTP). // Invalid Struct -> NOT 1|100
		Post("/api/class").
		JSON(`{"name":"Aerobics", "start_date":"2022-01-01", "end_date":"2022-01-03", "capacity":0 }`).
		Header("Content-Type", "application/json").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
