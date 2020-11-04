# Dockerized-Golang-Postgres-Mysql-API-with-Swaggo


## Introduction

this started as a learning ground to build a simple clean demo app in go that would act as a (mostly) rest api. I wanted a few things 


### Layout 

mostly created a MVC type of layout ( less V because we dont have any actual views )

### Dockerized

the app is dockerized and the test variant is dockerized too both app/test app are packaged with the .env var ( this is where the app reads the information from )
furtheremore the docker-compose files are setup to use the same variables ( just source .env to apply to your env )

### Databases

If you look inside the env file, you can define here which database to use: postgresql(default)/mysql/sqlite - sqlite is nice for local development if you do not want to run a db next to it as it stores to a flat file

We use gorm and gorm does not only create tables based on our model definitions, but also allows us to switch between the different database types easily

#### start main app
docker-compose up

#### start the app in testmode
docker-compose -f docker-compose.test.yml up


## Documentation

you want to know what endpoints are available and what params they need, checkout http://localhost:8080/swagger/index.html
you can update the swagger documentation my issuing a simple "swag init" inside the folder where the main.go is


## Authentication JWTs

/api/v1/user/login is an endpoint that accepts a json body with nickname/email/password and returns a Bearer token which is used as a header in subsequent requests that require an Authentication

the authentication part is implemented by using a api/middlewares/middlewares.go which allows easy setting of route specific rules inside the api/controllers/routes.go

## routing

This project uses mux as a router api/controllers/routes.go
