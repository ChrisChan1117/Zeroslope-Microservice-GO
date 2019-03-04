# Zeroslope GoLang
This is the zeroslope microservice architecture using GO. 

## Technologies
Here is a list of the technologies used in this project:
* We use GIN for http routing.
* We use VEGETA for load testing.
* We use GORM for our ORM layer.
* We use PG for postgres database access.

## Endpoints
Here is an overall layout of what endpoints come with this architecture:

| Method | Route       | Description                                  |
| ------ | ----------- | -------------------------------------------- |
| GET    | /health/    | Health check.                                |
| POST   | /auth/login | Creates a JWT token for access.              |
| GET    | /sample     | Gets a list of records.                      |
| GET    | /sample/:id | Gets a record by id.                         |
| POST   | /sample/    | Creates a record.                            |
| PUT    | /sample/    | Updates a record.                            |
| DELETE | /sample/:id | Deletes a record.                            |


## Development


### Installing dependencies
```bash
dep ensure
```

### Running locally
```
go run main.go
```

### Running in docker
```
docker-compose up
```
