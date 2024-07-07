# Gohex

[![Test Status](https://github.com/illud/gohex/actions/workflows/go.yml/badge.svg)](https://github.com/illud/gohex/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/illud/gohex?status.svg)](https://pkg.go.dev/github.com/illud/gohex?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/illud/gohex)](https://goreportcard.com/report/github.com/illud/gohex)
## Create project with Hexagonal Architecture folder structure inluding recommended way with Vertical Slicing

\
Gohex is a cli tool to create Hexagonal Architecture + Vertical Slicing app for you including gin-gonic, bcrypt and jwt.

- Creates Hexagonal Architecture + Vertical Slicing project for you.


## Features
- Hexagonal Architecture + Vertical Slicing Folder Structure (https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
- Gin Gonic (https://github.com/gin-gonic/gin)
- [Swagger](#swagger) (https://github.com/swaggo/gin-swagger)
- Jwt (https://github.com/dgrijalva/jwt-go)
- Bcrypt (https://golang.org/x/crypto/bcrypt)
- [Async](#async) - Async functions
- Auto add swagger for your endpoint
- [Modules](#modules) - Auto generate module with crud flow
- [Endpoint](#endpoint) - Auto add new endpoint(POST, GET, PUT, DELETE)
- [Database Service](#database-service) - Auto generate db service client 
  - Mysql
  - Gorm
- Example tasks api
- [Testing](#testing)  (Auto generate e2e test example when creating a new modules)

## Installation

Gohex requires [Go](https://golang.org/) v1.17 or later to run.

Install the dependencies.

```sh
go get github.com/illud/gohex
```
Or

```sh
go install github.com/illud/gohex@latest
```
## How to use

In your terminal type to see all avaible commands:

```sh
gohex
```

To create a new gin-gonic with Hexagonal Architecture + Vertical Slicing project(This includes a crud example with the name of Tasks):

```
▶  New Project
  Module
  Endpoint
  DB Service

Enter Project Name: yourprojectname
```

## Modules

To create a new module with crud flow:
please use snake_case when the module name consist of two or more words
```
  New Project
▶  Module
  Endpoint
  DB Service

Enter Module Name: your_module_name
```

## Endpoint
To add a new endpoint to your module:

```
  New Project
  Module
▶  Endpoint
  DB Service
```

Pick your module:
```
  example_module
  example_module_two
▶  current_module
```

Pick your method:
```
  POST
▶  GET
  PUT
  DELETE
```

## Database service
To create a new db service client with Mysql, Gorm or Prisma:

```
  New Project
  Module
  Endpoint
▶  DB Service
```

Mysql - to learn more visit (https://github.com/go-sql-driver/mysql)
```
Enter DB(mysql, gorm) Name: mysql
```

Gorm - to learn more visit (https://github.com/jinzhu/gorm)
```
Enter DB(mysql, gorm) Name: gorm
```

## This will generate a database connection in data/db.go

<br/>

### For Mysql and Gorm import the service in your repository like for example:

```go
db "yourProjectName/data"
```

Example for Mysql:
```go
// Insert new tasks
res, err := db.Client().Exec("INSERT INTO tasks VALUES(DEFAULT, 'Title', 'Desc')")
if err != nil {
  fmt.Println("ERROR: ", err)
}
fmt.Println(res)
```
To learn more visit (https://github.com/go-sql-driver/mysql)


Example for Gorm:
```go
// Insert new tasks
err := db.Client().Save(&tasksEntity.Task{
  Title:       "TEST",
  Description: "This is a description",
})

if err != nil {
  fmt.Println(err)
}
```
To learn more visit (https://github.com/jinzhu/gorm)

## Async
How to use async:

```go
package main

import async "yourProjectName/helpers"

//This functions wait 3 seconds to return 1
func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 1
}

func main() {
	fmt.Println("Let's start ...")

  //Here you call the function as async function
	future := async.Exec(func() interface{} {
		return DoneAsync()//The function that will await
	}).Await()

	fmt.Println("Done is running ...")
	fmt.Println(future)
}
```

# Swagger

Build your application and after that, go to http://localhost:5000/swagger/index.html , you to see your Swagger UI.

When you create a new module swagger will be added automatically then you only need to modified what you need, but remember each time you modified swagger use the next command 

```shell
swag init
```
To learn more visit (https://github.com/swaggo/gin-swagger)

<br/>

# Testing

To run tests use
```go
go test -v ./...
```
To get coverage
```go
go test -v -cover --coverprofile=coverage.out  -coverpkg=./... ./...
```
To view test coverage on your browser
```go
go tool cover -html=coverage.out
```
Total coverage

Windows
```go
go tool cover -func=coverage.out | findstr total:
```
Linux
```go
go tool cover -func=coverage.out | grep total:
```

Folder Structure:

```
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   main.go
│   README
│
├───adapters
│   ├───bcrypt
│   │       bcrypt.go
│   │
│   ├───database
│   │       db.go
│   │
│   └───jwt
│           jwt.go
│
├───app
│   └───tasks
│       ├───aplication
│       │       tasks.controller.go
│       │
│       ├───domain
│       │   ├───models
│       │   │       tasks.model.go
│       │   │
│       │   ├───repositories
│       │   │       tasks.repository.go        
│       │   │
│       │   └───services
│       │           tasks.service.go
│       │
│       └───infraestructure
│               tasks.db.go
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───e2e
│   └───tasks
│           gettasks_test.go
│
├───env
│       env.go
│
├───helpers
│       async.go
│       errors.go
│
└───router
        router.go
```

## Todo
- Add correct endpoint name when generating a new method

## License

MIT

Gohex is [MIT licensed](LICENSE).
