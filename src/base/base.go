package base

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func BaseData(folderName string) {
	//Add data to main.go
	mainString :=
		`package main

import (
	"fmt"
	"strconv"
	//Uncomment next line when you want to connect to a database
	//db "` + folderName + `/adapters/database"
	env "` + folderName + `/env"
	router "` + folderName + `/router"
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
	mainBytes := []byte(mainString)
	os.WriteFile(folderName+"/main.go", mainBytes, 0)

	//Add data to router.go
	routingString :=
		`package router

import (
	tasksController "` + folderName + `/app/tasks/aplication"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "` + folderName + `/docs"
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

	"github.com/gin-gonic/gin"
	tasksService "` + folderName + `/app/tasks/domain/services"
	tasksDatabase "` + folderName + `/app/tasks/infraestructure"
	tasksModel "` + folderName + `/app/tasks/domain/models"
)

// Create a TasksDb instance
var tasksDb = tasksDatabase.NewTasksDb()

// Create a Service instance using the TasksDb
var service = tasksService.NewService(tasksDb)

// Post Tasks
// @Summary Post Tasks
// @Schemes
// @Description Post Tasks
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body tasksModel.Task true "Body to create Tasks"
// @Success 200
// @Router /tasks [post]
func CreateTasks(c *gin.Context) {
	var task tasksModel.Task
	c.ShouldBindJSON(&task)

	c.JSON(200, gin.H{
		"data": service.CreateTasks(task.Id, task.Title, task.Description),
	})
}

// Get Tasks
// @Summary Get Tasks
// @Schemes
// @Description Get Tasks
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /tasks [Get]
func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": service.GetTasks(),
	})
}

// Get Tasks
// @Summary Get Tasks
// @Schemes
// @Description Get Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int64 true "taskId"
// @Accept json
// @Produce json
// @Success 200
// @Router /tasks/{taskId} [Get]
func GetOneTasks(c *gin.Context) {
	var task tasksModel.Task
	c.ShouldBindJSON(&task)

	taskId := c.Param("taskId")
	taskIdToInt64, _ := strconv.ParseInt(taskId, 10, 64)

	c.JSON(200, gin.H{
		"data": service.GetOneTasks(taskIdToInt64),
	})
}

// Put Tasks
// @Summary Put Tasks
// @Description Put Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int64 true "taskId"
// @Accept json
// @Produce json
// @Param Body body tasksModel.Task true "Body to update"
// @Success 200
// @Router /tasks/{taskId} [Put]
func UpdateTasks(c *gin.Context) {
	var task tasksModel.Task
	c.ShouldBindJSON(&task)
	taskId := c.Param("taskId")

	c.JSON(200, gin.H{
		"data": service.UpdateTasks(taskId, task.Title, task.Description),
	})
}

// Put Tasks
// @Summary Delete Tasks
// @Description Delete Tasks
// @Tags Tasks
// @Security BearerAuth
// @Param taskId path int64 true "taskId"
// @Accept json
// @Produce json
// @Success 200
// @Router /tasks/{taskId} [Delete]
func DeleteTasks(c *gin.Context) {
	taskId := c.Param("taskId")

	c.JSON(200, gin.H{
		"data": service.DeleteTasks(taskId),
	})
}`
	taskControllerBytes := []byte(taskControllerString)
	os.WriteFile(folderName+"/app/tasks/aplication/tasks.controller.go", taskControllerBytes, 0)

	//Add data to models/task.model.go
	taskModelString :=
		`package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
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
	tasksModel "` + folderName + `/app/tasks/domain/models"
)

type ITasks interface {
	CreateTasks(taskId int, title string, description string) string
	GetTasks() []*tasksModel.Task
	GetOneTasks(taskId int64) interface{}
	UpdateTasks(taskId string, title string, description string) string
	DeleteTasks(taskId string) string
}`
	taskRepositoryBytes := []byte(taskRepositoryString)
	os.WriteFile(folderName+"/app/tasks/domain/repositories/tasks.repository.go", taskRepositoryBytes, 0)

	//Add data to task.service.go
	taskServiceString :=
		`package services

import (
	tasksModel "` + folderName + `/app/tasks/domain/models"
	tasksInterface "` + folderName + `/app/tasks/domain/repositories"
)

type Service struct {
	tasksRepository tasksInterface.ITasks
}

func NewService(tasksRepository tasksInterface.ITasks) *Service {
	return &Service{
		tasksRepository: tasksRepository,
	}
}

func (s *Service) CreateTasks(taskId int, title string, description string) string {
	return s.tasksRepository.CreateTasks(taskId, title, description)
}

func (s *Service) GetTasks() []*tasksModel.Task {
	return s.tasksRepository.GetTasks()
}

func (s *Service) GetOneTasks(taskId int64) interface{} {
	return s.tasksRepository.GetOneTasks(taskId)
}

func (s *Service) UpdateTasks(taskId string, title string, description string) string {
	return s.tasksRepository.UpdateTasks(taskId, title, description)
}

func (s *Service) DeleteTasks(taskId string) string {
	return s.tasksRepository.DeleteTasks(taskId)
}`
	taskServiceBytes := []byte(taskServiceString)
	os.WriteFile(folderName+"/app/tasks/domain/services/tasks.service.go", taskServiceBytes, 0)

	//Add data to tasks.db.go
	taskInfraestructureString :=
		`package infraestructure

import (
	tasksModel "` + folderName + `/app/tasks/domain/models"
	// uncomment this a change _ for db when your are making database queries
	_ "` + folderName + `/adapters/database"
)

type TasksDb struct {
	// Add any dependencies or configurations related to the TasksRepository here, if needed.
}

func NewTasksDb() *TasksDb {
	// Initialize any dependencies and configurations for the TasksRepository here, if needed.
	return &TasksDb{}
}

func (t *TasksDb) CreateTasks(taskId int, title string, description string) string {
	return "Tasks created"
}

func (t *TasksDb) GetTasks() []*tasksModel.Task {
	var tasks []*tasksModel.Task 
	tasks = append(tasks, &tasksModel.Task{Id: 1, Title: "Hello", Description: "World"})
	return tasks
}

func (t *TasksDb) GetOneTasks(taskId int64) interface{} {
	return "one tasks"
}

func (t *TasksDb) UpdateTasks(taskId string, title string, description string) string {
	return "task updated"
}

func (t *TasksDb) DeleteTasks(taskId string) string {
	return "tasks deleted"
}`
	taskInfraestructureBytes := []byte(taskInfraestructureString)
	os.WriteFile(folderName+"/app/tasks/infraestructure/tasks.db.go", taskInfraestructureBytes, 0)

	//Add data to data/db.go
	taskDataString := `package data

import (
	"fmt"
	tasksModel "` + folderName + `/app/tasks/domain/models"

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
	router "` + folderName + `/router"
	token "` + folderName + `/adapters/jwt"

	/*
		- Uncomment this when you are testing real data coming from database.
		db "github.com/app/` + folderName + `/infraestructure"
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

	assert.Contains(t, w.Body.String(), expected, "🔴 Expected %v 🔴 got %v", expected, w.Body.String())
	assert.Contains(t, w.Result().Status, expectedStatus, "🔴 Expected %v 🔴 got %v", expectedStatus, w.Result().Status)
	fmt.Println("🟢")
}`

	tasksTestBytes := []byte(tasksTestString)
	os.WriteFile(folderName+"/e2e/tasks/getTasks_test.go", tasksTestBytes, 0)
}

func BaseModuleCrud(moduleName string, moduleNameSnakeCase string) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	//Add data to controller.go
	controllerString :=
		`package aplication

import (
	"strconv"

	"github.com/gin-gonic/gin"
	` + moduleName + `Services "` + currentDirName + `/app/` + moduleName + `/domain/services"
	` + moduleName + `Database "` + currentDirName + `/app/` + moduleName + `/infraestructure"
	// Replace for dto
	` + moduleName + `Model "` + currentDirName + `/app/` + moduleName + `/domain/models"
)

// Create a ` + moduleName + `Db instance
var ` + moduleName + `Db = ` + moduleName + `Database.New` + strings.Title(moduleName) + `Db()

// Create a Service instance using the ` + moduleName + `Db
var service = ` + moduleName + `Services.NewService(` + moduleName + `Db)

// Post ` + strings.Title(moduleName) + `
// @Summary Post ` + strings.Title(moduleName) + `
// @Schemes
// @Description Post ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body ` + moduleName + `Model.` + strings.Title(moduleName) + ` true "Body to create ` + strings.Title(moduleName) + `"
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + ` [Post]
func Create` + strings.Title(moduleName) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `
	c.ShouldBindJSON(&` + moduleName + `)

	c.JSON(200, gin.H{
		"data": service.Create` + strings.Title(moduleName) + `(),
	})
}

// Get ` + strings.Title(moduleName) + `
// @Summary Get ` + strings.Title(moduleName) + `
// @Schemes
// @Description Get ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + ` [Get]
func Get` + strings.Title(moduleName) + `(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": service.Get` + strings.Title(moduleName) + `(),
	})
}

// Get ` + strings.Title(moduleName) + `
// @Summary Get ` + strings.Title(moduleName) + `
// @Schemes
// @Description Get ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int64 true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Get]
func GetOne` + strings.Title(moduleName) + `(c *gin.Context) {
	` + moduleName + `Id := c.Param("` + moduleName + `Id")
	` + moduleName + `IdToInt64, _ := strconv.ParseInt(` + moduleName + `Id, 10, 64)

	c.JSON(200, gin.H{
		"data": service.GetOne` + strings.Title(moduleName) + `(` + moduleName + `IdToInt64),
	})
}

// Put ` + strings.Title(moduleName) + `
// @Summary Put ` + strings.Title(moduleName) + `
// @Schemes
// @Description Put ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int64 true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Param Body body ` + moduleName + `Model.` + strings.Title(moduleName) + ` true "Body to update ` + strings.Title(moduleName) + `"
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Put]
func Update` + strings.Title(moduleName) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `
	c.ShouldBindJSON(&` + moduleName + `)
	` + moduleName + `Id := c.Param("` + moduleName + `Id")
	` + moduleName + `IdToInt, _ := strconv.ParseInt(` + moduleName + `Id, 10, 64)

	c.JSON(200, gin.H{
		"data": service.Update` + strings.Title(moduleName) + `(` + moduleName + `IdToInt),
	})
}

// Delete ` + strings.Title(moduleName) + `
// @Summary Delete ` + strings.Title(moduleName) + `
// @Schemes
// @Description Delete ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int64 true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Delete]
func Delete` + strings.Title(moduleName) + `(c *gin.Context) {
	` + moduleName + `Id := c.Param("` + moduleName + `Id")
	` + moduleName + `IdToInt, _ := strconv.ParseInt(` + moduleName + `Id, 10, 64)

	c.JSON(200, gin.H{
		"data": service.Delete` + strings.Title(moduleName) + `(` + moduleName + `IdToInt),
	})
}`
	controllerBytes := []byte(controllerString)
	os.WriteFile("app/"+moduleName+"/aplication/"+moduleNameSnakeCase+".controller.go", controllerBytes, 0)

	//Add data to moduleName.modle.go
	modelString :=
		`package models

type ` + strings.Title(moduleName) + ` struct {
	Id    int
}`
	modelsBytes := []byte(modelString)
	os.WriteFile("app/"+moduleName+"/domain/models/"+moduleNameSnakeCase+".model.go", modelsBytes, 0)

	//Add data to usecase.go
	servicesString :=
		`package services

import (
	` + moduleName + `Model "` + currentDirName + `/app/` + moduleName + `/domain/models"
	` + moduleName + `Interface "` + currentDirName + `/app/` + moduleName + `/domain/repositories"
)

type Service struct {
	` + moduleName + `Repository ` + moduleName + `Interface.I` + strings.Title(moduleName) + `
}

func NewService(` + moduleName + `Repository ` + moduleName + `Interface.I` + strings.Title(moduleName) + `) *Service {
	return &Service{
		` + moduleName + `Repository: ` + moduleName + `Repository,
	}
}

func (s *Service) Create` + strings.Title(moduleName) + `() string {
	return s.` + moduleName + `Repository.Create` + strings.Title(moduleName) + `()
}

func (s *Service) Get` + strings.Title(moduleName) + `() []*` + moduleName + `Model.` + strings.Title(moduleName) + ` {
	return s.` + moduleName + `Repository.Get` + strings.Title(moduleName) + `()
}

func (s *Service) GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int64) interface{} {
	return s.` + moduleName + `Repository.GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id)
}

func (s *Service) Update` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string {
	return s.` + moduleName + `Repository.Update` + strings.Title(moduleName) + `(` + moduleName + `Id)
}

func (s *Service) Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string {
	return s.` + moduleName + `Repository.Delete` + strings.Title(moduleName) + `(` + moduleName + `Id)
}`
	servicesBytes := []byte(servicesString)
	os.WriteFile("app/"+moduleName+"/domain/services/"+moduleNameSnakeCase+".service.go", servicesBytes, 0)

	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`package repositories

import (
	` + moduleName + `Model "` + currentDirName + `/app/` + moduleName + `/domain/models"
)

type I` + strings.Title(moduleName) + ` interface {
	Create` + strings.Title(moduleName) + `() string
	Get` + strings.Title(moduleName) + `() []*` + moduleName + `Model.` + strings.Title(moduleName) + `
	GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int64) interface{}
	Update` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string
	Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string
}`
	repositoryInterfaceBytes := []byte(repositoryInterfaceString)
	os.WriteFile("app/"+moduleName+"/domain/repositories/"+moduleNameSnakeCase+".repository.go", repositoryInterfaceBytes, 0)

	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`package infraestructure

import (
	` + moduleName + `Model "` + currentDirName + `/app/` + moduleName + `/domain/models"
	// uncomment this a change _ for db when your are making database queries
	_ "` + currentDirName + `/adapters/database"
)

type ` + strings.Title(moduleName) + `Db struct {
	// Add any dependencies or configurations related to the UserRepository here, if needed.
}

func New` + strings.Title(moduleName) + `Db() *` + strings.Title(moduleName) + `Db {
	// Initialize any dependencies and configurations for the ` + strings.Title(moduleName) + `Repository here, if needed.
	return &` + strings.Title(moduleName) + `Db{}
}

var ` + moduleName + ` []` + moduleName + `Model.` + strings.Title(moduleName) + `


func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Create` + strings.Title(moduleName) + `() string {
	return "` + strings.Title(moduleName) + ` created"
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Get` + strings.Title(moduleName) + `() []*` + moduleName + `Model.` + strings.Title(moduleName) + ` {
	var ` + moduleName + ` []*` + moduleName + `Model.` + strings.Title(moduleName) + `
	` + moduleName + ` = append(` + moduleName + `, &` + moduleName + `Model.` + strings.Title(moduleName) + `{Id: 1})
	return ` + moduleName + `
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int64) interface{} {
	return ` + moduleName + `Id
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Update` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string {
	return "` + strings.Title(moduleName) + ` updated"
}

func  (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int64) string {
	return "` + strings.Title(moduleName) + ` deleted"
}`
	infraestructureBytes := []byte(infraestructureString)
	os.WriteFile("app/"+moduleName+"/infraestructure/"+moduleNameSnakeCase+".db.go", infraestructureBytes, 0)

	// TEST
	testString :=
		`package ` + moduleName + `_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	router "` + currentDirName + `/router"
	token "` + currentDirName + `/adapters/jwt"

	/*
		- Uncomment this when you are testing real data coming from database.
		db "github.com/app/` + currentDirName + `/infraestructure"
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

func TestGet` + strings.Title(moduleName) + `(t *testing.T) {
	// Call Setup/Teardown
	teardown := setup(t)
	defer teardown(t)

	tokenData := token.GenerateToken("test") //Your token data

	router := router.Router()

	w := httptest.NewRecorder()

	values := map[string]interface{}{"token": tokenData} // this is the body in case you make a post, put
	jsonValue, _ := json.Marshal(values)

	req, _ := http.NewRequest("GET", "/` + regex.StringToHyphen(moduleName) + `", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenData)

	// In case you use cookies like for example token
	req.Header.Set("Cookie", "token="+tokenData+";")

	router.ServeHTTP(w, req) 

	expected := ` + "`" + `{"data":[{"Id":1}]}` + "`" + ` // Your expected data inside backquote 
	expectedStatus := "200 OK"

	assert.Contains(t, w.Body.String(), expected, "🔴 Expected %v 🔴 got %v", expected, w.Body.String())
	assert.Contains(t, w.Result().Status, expectedStatus, "🔴 Expected %v 🔴 got %v", expectedStatus, w.Result().Status)
	fmt.Println("🟢")
}`
	//Add data to e2e
	testBytes := []byte(testString)
	os.WriteFile("e2e/"+moduleName+"/get_"+moduleNameSnakeCase+"_test.go", testBytes, 0)
}

func BaseDbClient(clientName string) {
	// Add database client
	clientString := ""
	if clientName == "mysql" {
		clientString =
			`package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// db The database connection
	db *sql.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := sql.Open("mysql", "databaseUsername:databasePassword@tcp(localhost:3306)/yourDatabaseTablename")

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	// defer db.Close()
	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *sql.DB {
	return db
}`
		//Adds db conection to main.go
		AppendDbConnectionToMain()
	}

	if clientName == "gorm" {
		clientString =
			`package database
		
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// db The database connection
	db *gorm.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := gorm.Open("mysql", "databaseUsername:databasePassword@tcp(127.0.0.1:3306)/yourDatabaseTablename?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *gorm.DB {
	return db
}`
		//Adds db conection to main.go
		AppendDbConnectionToMain()
	}

	// 	if clientName == "prisma" {

	// 		dir, err := os.Getwd()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(dir)
	// 		var ss []string
	// 		if runtime.GOOS == "windows" {
	// 			ss = strings.Split(dir, "\\")
	// 		} else {
	// 			ss = strings.Split(dir, "/")
	// 		}

	// 		currentDirName := ss[len(ss)-1]

	// 		clientString =
	// 			`package data

	// import (
	// 	"fmt"

	// 	"` + currentDirName + `/data/prisma/db"
	// 	"golang.org/x/net/context"
	// )

	// var (
	// 	// db The database connection
	// 	prismaDdb *db.PrismaClient
	// )

	// func Connect(){
	// 	client := db.NewClient()
	// 	if err := client.Prisma.Connect(); err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	// defer func() {
	// 	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 	// 		panic(err)
	// 	// 	}
	// 	// }()
	// 	prismaDdb = client
	// 	fmt.Println("CONNECTED")
	// }

	// func Client() *db.PrismaClient {
	// 	return prismaDdb
	// }

	// var Context = ContextService()

	// func ContextService() context.Context {
	// 	ctx := context.Background()
	// 	return ctx
	// }`

	// 		//Adds db conection to main.go
	// 		AppendDbConnectionToMain()

	// 		//Insertdata into prisma.schema
	// 		prismaString :=
	// 			`datasource db {
	// 	// could be postgresql or mysql
	// 	provider = "sqlite"
	// 	url      = "file:dev.db"
	// }

	// generator db {
	// 	provider = "go run github.com/prisma/prisma-client-go"
	// 	// set the output folder and package name
	// 	   output           = "./infraestructure/databases/prisma/db"
	// 	   package          = "db"
	// }

	// //This is and example table add your own schemas
	// model Tasks {
	// 	id        Int      @id @default(autoincrement())
	// 	createdAt DateTime @default(now())
	// 	updatedAt DateTime @updatedAt
	// 	title     String
	// 	description String
	// }`

	// 		prismaSChemaBytes := []byte(prismaString)
	// 		os.WriteFile("schema.prisma", prismaSChemaBytes, 0)
	// 	}

	//Add data to db.go
	clientBytes := []byte(clientString)
	os.WriteFile("/adapters/database/db.go", clientBytes, 0)
}

// ADD controller to router.go crud
func AppendToRoutingCrud(moduleName string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("router/router.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	` + moduleName + `Controller "` + currentDirName + `/app/` + moduleName + `/aplication"`
		}

		if strings.Contains(line, "return router") {
			lines[i] = ` //` + moduleName + `
	router.POST("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Create` + strings.Title(moduleName) + `)
	router.GET("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Get` + strings.Title(moduleName) + `)
	router.GET("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.GetOne` + strings.Title(moduleName) + `)
	router.PUT("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.Update` + strings.Title(moduleName) + `)
	router.DELETE("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.Delete` + strings.Title(moduleName) + `)

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt router/router.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt router/router.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}

// ADD controller to router.go simple
func AppendToRoutingSimple(moduleName string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("routing/routing.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	` + moduleName + `Controller "github.com/` + currentDirName + `/controller/` + moduleName + `"`
		}

		if strings.Contains(line, "return router") {
			lines[i] = ` //` + moduleName + `
	router.GET("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Get` + strings.Title(moduleName) + `)

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("routing/routing.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt routing/routing.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt routing/routing.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}

// ADD db conection to main.go
func AppendDbConnectionToMain() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	db "` + currentDirName + `/adapters/database"`
		}

		if strings.Contains(line, "router.Router().Run") {
			lines[i] = ` //Connect to database
			db.Connect()

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("main.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt main.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt main.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}
