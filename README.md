# glofox-task
Repository for glofox task

## How to run
Uses Docker & Docker-Compose. 

A [script](run.sh) is provided that uses docker-compose to bring the containers down and up every time.  
Note: It launches containers seperately since database was not being ready on-time.

## Project Structure
Based on the [Repository Pattern](https://dakaii.medium.com/repository-pattern-in-golang-d22d3fa76d91) which intends to asbtract the different layers.  
In this specific implementation of the repository pattern, the repository serves as a bridge between the database logic and the controller which due to the simplicity of the task also includes the business logic (Service).    

Flow: 
- User makes requests to Controller
- Controller calls Repository
- Repository calls specific Database implementation
- Database Implementation connects to Database

### Improvements:
Implementation of a Service layer that would abstract business logic from the repository creating the flow ```controller -> service -> repository -> dbImplementation -> db```


## Implementation Details
### Storage
For storing Classes and Bookings, I decided to use a simple Postgres instance through the use of an ORM, more specifically [GORM](https://gorm.io/).

### Endpoints
#### Bookings
Create Booking
```
curl -X POST http://localhost:8080/api/booking -H 'Content-Type: application/json' -d '{"name":"john", "date":"2012-04-23"}' 
```
Get all bookings
```
curl -X GET http://localhost:8080/api/bookings 
```
#### Classes
Create Class
```
curl -X POST http://localhost:8080/api/class -H 'Content-Type: application/json' -d '{"name":"Aerobics", "start_date":"2012-04-20", "end_date":"2012-04-25", "capacity":2 }' 
```
Get all classes
```
curl -X GET http://localhost:8080/api/class 
```

### Testing
For testing, I used the following [Framework](https://apitest.dev/).

Command to test   
Godotenv -> Initializes with .env     
go test ./path/file.go-> Executes that test file  
DATABASE_HOST -> Overwrites host for localhost so we can access docker container  
```
DATABASE_HOST=localhost godotenv -f .env go test ./path/file.go
```

It stands to note that we must first execute the Class tests before the Booking tests since for success bookings we require a Class to exist. 

### Documentation
For the documentation of the application,an implementation of Swagger, [Swag](https://github.com/swaggo/swag#the-swag-formatter) was used.

The link for accessing the swagger page once it is up and running is [http://localhost:8080/swagger/](http://localhost:8080/swagger/index.html)

### Others
Decoding JSON body based on: [blog](https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body)