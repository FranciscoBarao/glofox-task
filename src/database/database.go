package database

import (
	"errors"
	"fmt"
	"glofox-task/middleware"
	"glofox-task/models"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresqlRepository struct {
	db *gorm.DB
}

func Connect() (*PostgresqlRepository, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting to the database: " + err.Error())
		return nil, err
	}

	log.Println("Connected to the Database")

	db.AutoMigrate(models.Booking{})
	db.AutoMigrate(models.Class{})

	log.Println("Database Migrations complete")

	return &PostgresqlRepository{db}, nil
}

func getConfig() (string, error) {
	log.Println("Fetching env vars for Database")

	host, hostPresent := os.LookupEnv("DATABASE_HOST")
	user, userPresent := os.LookupEnv("POSTGRES_USER")
	pass, passPresent := os.LookupEnv("POSTGRES_PASSWORD")
	dbname, dbnamePresent := os.LookupEnv("POSTGRES_DB")
	port, portPresent := os.LookupEnv("DATABASE_PORT")

	if !hostPresent || !userPresent || !passPresent || !dbnamePresent || !portPresent {
		log.Println("Error occurred while fetching env vars")
		return "", errors.New("error occurred while fetching env vars")
	}

	log.Println("host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port)
	return "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port, nil
}

func (instance *PostgresqlRepository) Create(model interface{}) error {

	result := instance.db.Create(model)
	if result.Error != nil {
		log.Println("Error while creating a database entry: " + fmt.Sprintf("%v", model))
		if errors.Is(result.Error, gorm.ErrRegistered) {
			return middleware.NewCustomError(http.StatusConflict, "Entry already registered")
		}
		return result.Error
	}

	log.Println("Created database entry: " + fmt.Sprintf("%v", model))
	return nil
}

func (instance *PostgresqlRepository) ReadAll(model interface{}, condition, value string) error {

	result := instance.db.Find(model)
	if result.Error != nil {
		log.Println("Error while creating a database entry: " + fmt.Sprintf("%v", model))
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error Record not found: " + condition + " " + value)
			return middleware.NewCustomError(http.StatusNotFound, "Record Not found")
		}
		return result.Error
	}

	log.Println("Created database entry: " + fmt.Sprintf("%v", model))
	return nil
}
