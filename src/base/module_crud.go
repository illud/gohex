package base

import (
	"log"
	"os"
	"runtime"
	"strings"

	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

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
	"net/http"
	"github.com/gin-gonic/gin"
	` + moduleName + `Services "github.com/` + currentDirName + `/app/` + moduleName + `/domain/services"
	` + moduleName + `Database "github.com/` + currentDirName + `/app/` + moduleName + `/infraestructure"
	// Replace for dto
	` + moduleName + `Model "github.com/` + currentDirName + `/app/` + moduleName + `/domain/models"
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
	if err := c.ShouldBindJSON(&` + moduleName + `); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Create` + strings.Title(moduleName) + `(` + moduleName + `)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "` + moduleName + ` created",
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
	result, err := service.Get` + strings.Title(moduleName) + `()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Get ` + strings.Title(moduleName) + `
// @Summary Get ` + strings.Title(moduleName) + `
// @Schemes
// @Description Get ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Get]
func GetOne` + strings.Title(moduleName) + `(c *gin.Context) {
	` + moduleName + `Id, err := strconv.Atoi(c.Param("` + moduleName + `Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Put ` + strings.Title(moduleName) + `
// @Summary Put ` + strings.Title(moduleName) + `
// @Schemes
// @Description Put ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Param Body body ` + moduleName + `Model.` + strings.Title(moduleName) + ` true "Body to update ` + strings.Title(moduleName) + `"
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Put]
func Update` + strings.Title(moduleName) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `
	if err := c.ShouldBindJSON(&` + moduleName + `); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	` + moduleName + `Id, err := strconv.Atoi(c.Param("` + moduleName + `Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	 err = service.Update` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "` + moduleName + ` updated",
	})
}

// Delete ` + strings.Title(moduleName) + `
// @Summary Delete ` + strings.Title(moduleName) + `
// @Schemes
// @Description Delete ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + moduleName + `Id path int true "` + strings.Title(moduleName) + `Id"
// @Accept json
// @Produce json
// @Success 200
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Delete]
func Delete` + strings.Title(moduleName) + `(c *gin.Context) {
	` + moduleName + `Id, err := strconv.Atoi(c.Param("` + moduleName + `Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.Delete` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "` + moduleName + ` deleted",
	})
}
`
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
	` + moduleName + `Model "github.com/` + currentDirName + `/app/` + moduleName + `/domain/models"
	` + moduleName + `Interface "github.com/` + currentDirName + `/app/` + moduleName + `/domain/repositories"
)

type Service struct {
	` + moduleName + `Repository ` + moduleName + `Interface.I` + strings.Title(moduleName) + `
}

func NewService(` + moduleName + `Repository ` + moduleName + `Interface.I` + strings.Title(moduleName) + `) *Service {
	return &Service{
		` + moduleName + `Repository: ` + moduleName + `Repository,
	}
}

func (s *Service) Create` + strings.Title(moduleName) + `(` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `)  error {
	err := s.` + moduleName + `Repository.Create` + strings.Title(moduleName) + `(` + moduleName + `)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get` + strings.Title(moduleName) + `() ([]` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	result, err := s.` + moduleName + `Repository.Get` + strings.Title(moduleName) + `()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int) (` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	result, err := s.` + moduleName + `Repository.GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		return ` + moduleName + `Model.` + strings.Title(moduleName) + `{} , err
	}
	return result, nil
}

func (s *Service) Update` + strings.Title(moduleName) + `(` + moduleName + `Id int) error {
	err := s.` + moduleName + `Repository.Update` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int) error {
	err := s.` + moduleName + `Repository.Delete` + strings.Title(moduleName) + `(` + moduleName + `Id)
	if err != nil {
		return err
	}
	return nil
}`
	servicesBytes := []byte(servicesString)
	os.WriteFile("app/"+moduleName+"/domain/services/"+moduleNameSnakeCase+".service.go", servicesBytes, 0)

	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`package repositories

import (
	` + moduleName + `Model "github.com/` + currentDirName + `/app/` + moduleName + `/domain/models"
)

type I` + strings.Title(moduleName) + ` interface {
	Create` + strings.Title(moduleName) + `(` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `) error
	Get` + strings.Title(moduleName) + `() ([]` + moduleName + `Model.` + strings.Title(moduleName) + `, error)
	GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int) (` + moduleName + `Model.` + strings.Title(moduleName) + `, error)
	Update` + strings.Title(moduleName) + `(` + moduleName + `Id int) error
	Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int) error
}`
	repositoryInterfaceBytes := []byte(repositoryInterfaceString)
	os.WriteFile("app/"+moduleName+"/domain/repositories/"+moduleNameSnakeCase+".repository.go", repositoryInterfaceBytes, 0)

	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`package infraestructure

import (
	` + moduleName + `Model "github.com/` + currentDirName + `/app/` + moduleName + `/domain/models"
	// uncomment this a change _ for db when you are making database queries
	_ "github.com/` + currentDirName + `/adapters/database"
)

type ` + strings.Title(moduleName) + `Db struct {
	// Add any dependencies or configurations related to the UserRepository here if needed.
}

func New` + strings.Title(moduleName) + `Db() *` + strings.Title(moduleName) + `Db {
	// Initialize any dependencies and configurations for the ` + strings.Title(moduleName) + `Repository here if needed.
	return &` + strings.Title(moduleName) + `Db{}
}

var ` + moduleName + ` []` + moduleName + `Model.` + strings.Title(moduleName) + `

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Create` + strings.Title(moduleName) + `(` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `) error {
	// Implement your creation logic here
	return nil
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) Get` + strings.Title(moduleName) + `() ([]` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	// Implement your retrieval logic here
	var ` + moduleName + ` []` + moduleName + `Model.` + strings.Title(moduleName) + `
	` + moduleName + ` = append(` + moduleName + `, ` + moduleName + `Model.` + strings.Title(moduleName) + `{Id: 1})
	return ` + moduleName + `, nil
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) GetOne` + strings.Title(moduleName) + `(` + moduleName + `Id int) (` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	// Implement your single retrieval logic here
	return ` + moduleName + `Model.` + strings.Title(moduleName) + `{Id: ` + moduleName + `Id}, nil
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` ` + strings.Title(moduleName) + `Db) Update` + strings.Title(moduleName) + `(` + moduleName + `Id int)  error {
	// Implement your update logic here
	return nil
}

func (` + str.GetFirstCharacterOfString(moduleName) + ` ` + strings.Title(moduleName) + `Db) Delete` + strings.Title(moduleName) + `(` + moduleName + `Id int) error {
	// Implement your deletion logic here
	return nil
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
	router "github.com/` + currentDirName + `/router"
	token "github.com/` + currentDirName + `/adapters/jwt"

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
