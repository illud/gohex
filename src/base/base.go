package base

import (
	"os"
)

func BaseData(folderName string) {
	//Add data to tracker.json
	trackerString :=
		`{
  "modules": [{
    "moduleName": "tasks",
    "endpointName": "tasks"
  }]
}`
	mainBytes := []byte(trackerString)
	os.WriteFile(folderName+"/tracker.json", mainBytes, 0)

	//Add data to main.go
	mainString :=
		`package main

import (
	"fmt"
	"strconv"
	//Uncomment next line when you want to connect to a database
	//db "github.com/` + folderName + `/adapters/database"
	env "github.com/` + folderName + `/env"
	router "github.com/` + folderName + `/router"
)

//The next lines are for swagger docs
// @title ` + folderName + `
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	//Uncomment next line when you want to connect to a database
	//Connect to database
	//db.Connect()

	//Load .env port
	port := strconv.Itoa(env.Load().PORT)

	if port == "" {
		fmt.Println("$PORT must be set")
	}

	router.Router().Run(":" + port)
}`
	mainBytes = []byte(mainString)
	os.WriteFile(folderName+"/main.go", mainBytes, 0)

	//Add data to router.go
	routingString :=
		`package router

import (
	tasksController "github.com/` + folderName + `/app/tasks/aplication"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "github.com/` + folderName + `/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	//this sets gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.Default())

	//SWAGGER
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.POST("/tasks", tasksController.CreateTasks)
	router.GET("/tasks", tasksController.GetTasks)
	router.GET("/tasks/:taskId", tasksController.GetOneTasks)
	router.PUT("/tasks/:taskId", tasksController.UpdateTasks)
	router.DELETE("/tasks/:taskId", tasksController.DeleteTasks)

	return router
}`
	routingBytes := []byte(routingString)
	os.WriteFile(folderName+"/router/router.go", routingBytes, 0)

	//Add data to .env
	dotEnvString :=
		`PORT = 5000

VERSION = 1.0.0`

	dotEnvBytes := []byte(dotEnvString)
	os.WriteFile(folderName+"/.env", dotEnvBytes, 0)

	//Add data to .gitignore
	gitignoreString :=
		`# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# Env
*.env
*.env.test

# Tmp folder
tmp/`

	gitignoreBytes := []byte(gitignoreString)
	os.WriteFile(folderName+"/.gitignore", gitignoreBytes, 0)

	//Add data to README
	readmeString :=
		`
		_____       _               
	   |  __ \     | |              
	   | |  \/ ___ | |__   _____  __
	   | | __ / _ \| '_ \ / _ \ \/ /
	   | |_\ \ (_) | | | |  __/>  < 
		\____/\___/|_| |_|\___/_/\_\
   
[circleci-image]: https://img.shields.io/circleci/build/github/nestjs/nest/master?token=abc123def456
[circleci-url]: https://circleci.com/gh/nestjs/nest

<p align="center">A progressive <a href="http://golang.dev" target="_blank">Go</a> framework for building efficient and scalable server-side applications.</p>
<p align="center">

[![Test Status](https://github.com/illud/gohex/actions/workflows/go.yml/badge.svg)](https://github.com/illud/gohex/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/illud/gohex?status.svg)](https://pkg.go.dev/github.com/illud/gohex?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/illud/gohex)](https://goreportcard.com/report/github.com/illud/gohex)

</p>
<!--[![Backers on Open Collective](https://opencollective.com/nest/backers/badge.svg)](https://opencollective.com/nest#backer)
[![Sponsors on Open Collective](https://opencollective.com/nest/sponsors/badge.svg)](https://opencollective.com/nest#sponsor)-->

## Description

[Gohex](https://github.com/illud/gohex) CLI tool to create Hexagonal Architecture + Vertical Slicing.

## Running the app

` + "```" + `bash
# development
$ go run main.go

# watch mode
# for more go to https://github.com/gravityblast/fresh
$ fresh

# production mode
$ go build main.go
` + "```" + `

## Test

` + "```" + `bash
# tests
$ go test -v ./...

# to get coverage
$ go test -v -cover --coverprofile=coverage.out  -coverpkg=./... ./...

# to view test coverage on your browser
$ go tool cover -html=coverage.out

# prints formatted test output, and a summary of the test run
# for more go to https://github.com/gotestyourself/gotestsum
$ gotestsum --format testname
` + "```" + `

## Support

Gohex is an MIT-licensed open source project.

## Stay in touch

- Author - [Illud](https://github.com/illud)

## License

Nest is [MIT licensed](LICENSE).
"# Gohex" 
`

	readmeBytes := []byte(readmeString)
	os.WriteFile(folderName+"/README", readmeBytes, 0)

	//Add data to env.go
	envString :=
		`package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//Add here your .env data
type Env struct {
	PORT                   int
	VERSION                string
}

func Load() Env {
	godotenv.Load() //This loads your .env

	//Converts port string to int
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	//Returns .env data int Env struct
	return Env{
		PORT:                   port,
		VERSION:      os.Getenv("VERSION"),
	}
}`

	envBytes := []byte(envString)
	os.WriteFile(folderName+"/env/env.go", envBytes, 0)

	//Add data to task-controller.go
	taskControllerString :=
		`package tasks

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	tasksService "github.com/` + folderName + `/app/tasks/domain/services"
	tasksDatabase "github.com/` + folderName + `/app/tasks/infraestructure"
	tasksModel "github.com/` + folderName + `/app/tasks/domain/models"
)

// Create a TasksDb instance
var tasksDb = tasksDatabase.NewTasksDb()

// Create a Service instance using the TasksDb
var service = tasksService.NewService(tasksDb)

// CreateTasks handles the creation of a new task.
// @Summary Create Tasks
// @Description Create Tasks
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body tasksModel.Task true "Body to create Tasks"
// @Success 200 {string} string "Task created successfully"
// @Router /tasks [post]
func CreateTasks(c *gin.Context) {
	var task tasksModel.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateTasks(task.Id, task.Title, task.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "task created"})
}

// GetTasks retrieves all tasks.
// @Summary Get Tasks
// @Description Get Tasks
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} tasksModel.Task[]
// @Router /tasks [Get]
func GetTasks(c *gin.Context) {
	result, err := service.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetOneTasks retrieves a single task by taskId.
// @Summary Get Tasks
// @Description Get Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int true "taskId"
// @Accept json
// @Produce json
// @Success 200 {object} tasksModel.Task
// @Router /tasks/{taskId} [Get]
func GetOneTasks(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid taskId"})
		return
	}

	result, err := service.GetOneTasks(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// UpdateTasks handles updating an existing task by taskId.
// @Summary Update Tasks
// @Description Update Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int true "taskId"
// @Accept json
// @Produce json
// @Param Body body tasksModel.Task true "Body to update"
// @Success 200 {string} string "Task updated successfully"
// @Router /tasks/{taskId} [Put]
func UpdateTasks(c *gin.Context) {
	taskId := c.Param("taskId")
	var task tasksModel.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.UpdateTasks(taskId, task.Title, task.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "task updated"})
}

// DeleteTasks handles deleting a task by taskId.
// @Summary Delete Tasks
// @Description Delete Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int true "taskId"
// @Accept json
// @Produce json
// @Success 200 {string} string "Task deleted successfully"
// @Router /tasks/{taskId} [Delete]
func DeleteTasks(c *gin.Context) {
	taskId := c.Param("taskId")

	err := service.DeleteTasks(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "task deleted"})
}`
	taskControllerBytes := []byte(taskControllerString)
	os.WriteFile(folderName+"/app/tasks/aplication/tasks.controller.go", taskControllerBytes, 0)

	//Add data to models/task.model.go
	taskModelString :=
		`package models

type Task struct {
	Id          int
	Title       string
	Description string
}`
	taskModelBytes := []byte(taskModelString)
	os.WriteFile(folderName+"/app/tasks/domain/models/tasks.model.go", taskModelBytes, 0)

	//Add data to repositories/tasks.repository.go
	taskRepositoryString :=
		`package repositories

import (
	tasksModel "github.com/` + folderName + `/app/tasks/domain/models"
)

type ITasks interface {
	CreateTasks(taskId int, title string, description string)  error
	GetTasks() ([]tasksModel.Task, error)
	GetOneTasks(taskId int) (tasksModel.Task, error)
	UpdateTasks(taskId string, title string, description string)  error
	DeleteTasks(taskId string) error
}`
	taskRepositoryBytes := []byte(taskRepositoryString)
	os.WriteFile(folderName+"/app/tasks/domain/repositories/tasks.repository.go", taskRepositoryBytes, 0)

	//Add data to task.service.go
	taskServiceString :=
		`package services

import (
	tasksModel "github.com/` + folderName + `/app/tasks/domain/models"
	tasksInterface "github.com/` + folderName + `/app/tasks/domain/repositories"
)

type Service struct {
	tasksRepository tasksInterface.ITasks
}

func NewService(tasksRepository tasksInterface.ITasks) *Service {
	return &Service{
		tasksRepository: tasksRepository,
	}
}

func (s *Service) CreateTasks(taskId int, title string, description string) error {
	 err := s.tasksRepository.CreateTasks(taskId, title, description)
	if err != nil {
		return  err
	}
	return nil
}

func (s *Service) GetTasks() ([]tasksModel.Task, error) {
	tasks, err := s.tasksRepository.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Service) GetOneTasks(taskId int) (tasksModel.Task, error) {
	task, err := s.tasksRepository.GetOneTasks(taskId)
	if err != nil {
		return tasksModel.Task{}, err
	}
	return task, nil
}

func (s *Service) UpdateTasks(taskId string, title string, description string) error {
	err := s.tasksRepository.UpdateTasks(taskId, title, description)
	if err != nil {
		return  err
	}
	return  nil
}

func (s *Service) DeleteTasks(taskId string) error {
	err := s.tasksRepository.DeleteTasks(taskId)
	if err != nil {
		return  err
	}
	return  nil
}`
	taskServiceBytes := []byte(taskServiceString)
	os.WriteFile(folderName+"/app/tasks/domain/services/tasks.service.go", taskServiceBytes, 0)

	//Add data to tasks.db.go
	taskInfraestructureString :=
		`package infraestructure

import (
	tasksModel "github.com/` + folderName + `/app/tasks/domain/models"
	// uncomment this a change _ for db when your are making database queries
	_ "github.com/` + folderName + `/adapters/database"
	"errors"
)

type TasksDb struct {
	// Add any dependencies or configurations related to the TasksRepository here, if needed.
}

func NewTasksDb() *TasksDb {
	// Initialize any dependencies and configurations for the TasksRepository here, if needed.
	return &TasksDb{}
}

func (t *TasksDb) CreateTasks(taskId int, title string, description string) error {
	if taskId == 0 {
		return errors.New("taskId is required")
	}

	var task tasksModel.Task
	task.Id = 1
	task.Title = title
	task.Description = description

	return  nil
}

func (t *TasksDb) GetTasks() ([]tasksModel.Task, error) {
	var tasks []tasksModel.Task
	tasks = append(tasks, tasksModel.Task{Id: 1, Title: "Hello", Description: "World"})
	return tasks, nil
}

func (t *TasksDb) GetOneTasks(taskId int) (tasksModel.Task, error) {
	task := tasksModel.Task{Id: taskId, Title: "Sample Task", Description: "Sample Description"}
    if task.Id == 0 {
        return tasksModel.Task{}, errors.New("Task not found")
    }
	return task, nil
}

func (t *TasksDb) UpdateTasks(taskId string, title string, description string) error {
	if taskId == "" {
		return  errors.New("taskId is required")
	}
	return nil
}

func (t *TasksDb) DeleteTasks(taskId string)  error {
	if taskId == "" {
		return errors.New("taskId is required")
	}
	return nil
}`
	taskInfraestructureBytes := []byte(taskInfraestructureString)
	os.WriteFile(folderName+"/app/tasks/infraestructure/tasks.db.go", taskInfraestructureBytes, 0)

	//Add data to data/db.go
	taskDataString := `package data

import (
	"fmt"
	tasksModel "github.com/` + folderName + `/app/tasks/domain/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// DB The database connection
	db *gorm.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	// Migrate the schema
	dbCon.AutoMigrate(&tasksModel.Task{})

	// defer db.Close()
	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *gorm.DB {
	return db
}`

	taskDataBytes := []byte(taskDataString)
	os.WriteFile(folderName+"/adapters/database/db.go", taskDataBytes, 0)

	//Add data to adapters/bcrypt/bcrypt.go
	bcrypt := `package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

//hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

//check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}`

	bcryptBytes := []byte(bcrypt)
	os.WriteFile(folderName+"/adapters/bcrypt/bcrypt.go", bcryptBytes, 0)

	// ASYNC to helpers
	asyncString :=
		`package helpers

import "context"

// Future interface has the method signature for await
type Future interface {
	Await() interface{}
}

type future struct {
	await func(ctx context.Context) interface{}
}

func (f future) Await() interface{} {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func() interface{}) Future {
	var result interface{}
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = f()
	}()
	return future{
		await: func(ctx context.Context) interface{} {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return result
			}
		},
	}
}`
	//Add data to async.go
	asyncBytes := []byte(asyncString)
	os.WriteFile(folderName+"/helpers/async.go", asyncBytes, 0)

	// jwt
	jwtString :=
		`package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func ValidateToken(validate string) string {
	var tokenCheker string
	tokenString := validate
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println("Error: ", err)
		tokenCheker = "Error"
	} else {
		tokenCheker = "Ok"
	}

	return tokenCheker
}`
	//Add data to jwt.go
	jwtBytes := []byte(jwtString)
	os.WriteFile(folderName+"/adapters/jwt/jwt.go", jwtBytes, 0)

	// ERRORS
	errorsString :=
		`package helpers
import (
	"encoding/json"
)

type Error struct{
	Error string 
	Code int 
}

func ErrorJson(error string, code int) string {
	jsondata := &Error{error, code}
	encodejson, _ := json.Marshal(jsondata)
	return string(encodejson)
}
	
var BadRequest = ErrorJson("Bad Request", 400)
var Forbidden = ErrorJson("Forbidden", 403)
var NotFound = ErrorJson("Not Found", 404)
var Unauthorized = ErrorJson("Unauthorized", 401)`

	//Add data to errors.go
	errorsBytes := []byte(errorsString)
	os.WriteFile(folderName+"/helpers/errors.go", errorsBytes, 0)

	// getTasks_test.go
	tasksTestString :=
		`package tasks_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	router "github.com/` + folderName + `/router"
	token "github.com/` + folderName + `/adapters/jwt"

	/*
		- Uncomment this when you are testing real data coming from database.
		db "github.com/app/` + folderName + `/adapters/database"
	*/
)

// Setup and Teardown
func setup(t *testing.T) func(t *testing.T) {
	// Setup
	t.Log("setup sub test")

	// For test db
	t.Setenv("ENV", "TEST")

	/*
		- Uncomment this when you are testing real data coming from test database.
		db.Connect()
	*/

	// Teardown
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func TestGetTasks(t *testing.T) {
	// Call Setup/Teardown
	teardown := setup(t)
	defer teardown(t)

	tokenData := token.GenerateToken("test") //Your token data

	router := router.Router()

	w := httptest.NewRecorder()

	values := map[string]interface{}{"token": tokenData} // this is the body in case you make a post, put
	jsonValue, _ := json.Marshal(values)

	req, _ := http.NewRequest("GET", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenData)

	// In case you use cookies like for example token
	req.Header.Set("Cookie", "token="+tokenData+";")

	router.ServeHTTP(w, req)

	expected := ` + "`" + `{"data":[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Id":1,"Title":"Hello","Description":"World"}]}` + "`" + ` // Your expected data inside backquote
	expectedStatus := "200 OK"

	assert.Contains(t, w.Body.String(), expected, "Expected %v got %v", expected, w.Body.String())
	assert.Contains(t, w.Result().Status, expectedStatus, "Expected %v got %v", expectedStatus, w.Result().Status)
}`

	tasksTestBytes := []byte(tasksTestString)
	os.WriteFile(folderName+"/e2e/tasks/getTasks_test.go", tasksTestBytes, 0)
}
