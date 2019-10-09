# Zeroslope GoLang
This is the zeroslope microservice architecture using GO. Average response times are around ~2ms for api calls that use a database, and non-database api calls ~28μs. Memory footprint seems to be around 9.9mb of ram allocated to the go process when running the most expensive call.

## Technologies
Here is a list of the technologies used in this project:
* We use GLIDE for package management.
* We use SWAG for swagger generation.
* We use DLV for Debugging in VSCODE
* We use GIN for http routing.
* We use VEGETA for load testing.
* We use GORM for our ORM layer.
* We use PG for postgres database access.

## Endpoints
Here is an overall layout of what endpoints come with this architecture:

| Method | Route                  | Description                                  |
| ------ | ---------------------- | -------------------------------------------- |
| N/A    | /swagger/index.html    | Swagger UI                                   |
| GET    | /health/               | Health check.                                |
| POST   | /auth/login            | Creates a JWT token for access.              |
| GET    | /sample                | Gets a list of records.                      |
| GET    | /sample/:id            | Gets a record by id.                         |
| POST   | /sample/               | Creates a record.                            |
| PUT    | /sample/               | Updates a record.                            |
| DELETE | /sample/:id            | Deletes a record.                            |


## Development
Currently I write my golang in VSCODE and use the following commands:

### Installing dependencies
```bash
glide install
```

### Updating swagger docs
```bash
cd src; swag init; cd ..
```

### Running locally
```bash
go run main.go
```
Visit: `http://localhost:8080/swagger/index.html`

### Setting up postgres in docker
```bash
docker run --name pgdb -p 5432:5432 -e POSTGRES_PASSWORD=P4ssw0rd! -e POSTGRES_DB=zeroslope -d postgres
```

## Golang Dev Environment
TBD