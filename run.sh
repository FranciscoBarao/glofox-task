# Clear Console
printf "\033c"

# Closes running services
docker-compose down

# Generates Swag files
export PATH=$(go env GOPATH)/bin:$PATH # Required for swag init
( cd src ; swag init --parseDependency --parseInternal )

# Brings database service up
docker-compose up -d --build glofox-task-db

sleep 3 # Database was not ready when gorm attempted to connect

# Testing - Booking tests require class_test.go to run first since a class must exist for a booking to be created. Tests also leave the database with some dummy data
( cd src ; DATABASE_HOST=localhost godotenv -f .env go test ./tests/class_test.go )
( cd src ; DATABASE_HOST=localhost godotenv -f .env go test ./tests/booking_test.go ) 

# Brings server service up
docker-compose up -d --build glofox-task

#Docker logging
docker logs --follow glofox-task