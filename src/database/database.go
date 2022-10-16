package database

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"glofox-task/middleware"
	"glofox-task/models"
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

func (instance *PostgresqlRepository) ReadAll(model interface{}) error {

	result := instance.db.Preload(clause.Associations).Find(model)
	if result.Error != nil {
		log.Println("Error while Reading database entries: " + result.Error.Error())
		return result.Error
	}

	log.Println("Fetched all entries: " + fmt.Sprintf("%v", model))
	return nil
}

func (instance *PostgresqlRepository) ReadByCondition(value interface{}, condition string, variables ...interface{}) error {

	result := instance.db.Where(condition, variables...).First(value)

	if result.Error != nil {
		log.Println("Error while reading a database entry using condition: " + condition)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error Record not found with condition: " + condition + " and variables: " + fmt.Sprintf("%v", variables...))

			return middleware.NewCustomError(http.StatusNotFound, "Record not found")
		}
		return result.Error
	}

	log.Println("Fetched database entry: " + fmt.Sprintf("%v", value))
	return nil
}

func (instance *PostgresqlRepository) Count(value interface{}, count *int64, condition string, variables ...interface{}) error {

	result := instance.db.Model(value).Where(condition, variables...).Count(count)

	if result.Error != nil {
		log.Println("Error while counting a database entry using condition: " + condition)
		return result.Error
	}

	log.Println("Counted database entries: " + fmt.Sprintf("%v", value) + " count: " + fmt.Sprintf("%d", count))
	return nil
}

func (instance *PostgresqlRepository) Delete(value interface{}) error {

	result := instance.db.Delete(value)
	if result.Error != nil {
		log.Println("Error while deleting a database entry: " + fmt.Sprintf("%v", value))
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return middleware.NewCustomError(http.StatusNotFound, "Record Not found")
		}
		return result.Error
	}

	log.Println("Deleted database entry: " + fmt.Sprintf("%v", value))
	return nil
}
