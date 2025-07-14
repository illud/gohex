package base

import (
	"log"
	"os"
	"strings"

	append "github.com/illud/gohex/src/utils/append"
	utils "github.com/illud/gohex/src/utils/append"
	find "github.com/illud/gohex/src/utils/find"
	str "github.com/illud/gohex/src/utils/strings"
)

func GetMethod(moduleName string, methodName string) {
	trackerResult := utils.ReadTrackerFile()
	var endpointName string

	for _, module := range trackerResult.Modules {
		if module.ModuleName == moduleName {
			endpointName = module.EndpointName
			break
		}
	}

	//Add data to controller.go
	controllerString :=
		`
// Get ` + caser.String(moduleName) + `
// @Summary Get ` + caser.String(moduleName) + `
// @Schemes
// @Description Get ` + caser.String(moduleName) + `
// @Tags ` + caser.String(moduleName) + `
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /` + endpointName + `/` + strings.ToLower(methodName) + ` [Get]
func ` + str.DashToCamel(methodName) + `(c *gin.Context) {
	result, err := service.` + str.DashToCamel(methodName) + `()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
`
	controllerResult, err := find.FindFile("app/" + moduleName + "/aplication/")
	if err != nil {
		log.Fatal(err)
	}
	// Write the data to the end of the file
	append.AppendDataToFile("app/"+moduleName+"/aplication/"+*controllerResult, controllerString)

	// 	//Add data to service.go
	servicesString :=
		`
func (s *Service) ` + str.DashToCamel(methodName) + `() ([]` + moduleName + `Model.` + caser.String(moduleName) + `, error) {
	result, err := s.` + moduleName + `Repository.` + str.DashToCamel(methodName) + `()
	if err != nil {
		return nil, err
	}
	return result, nil
}`
	serviceResult, err := find.FindFile("app/" + moduleName + "/domain/services/")
	if err != nil {
		log.Fatal(err)
	}
	append.AppendDataToFile("app/"+moduleName+"/domain/services/"+*serviceResult, servicesString)

	// 	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`	` + str.DashToCamel(methodName) + `() ([]` + moduleName + `Model.` + caser.String(moduleName) + `, error)
}`

	repositoryResult, err := find.FindFile("app/" + moduleName + "/domain/repositories/")
	if err != nil {
		log.Fatal(err)
	}
	err = append.ReplaceLastCharacter("app/"+moduleName+"/domain/repositories/"+*repositoryResult, "}", repositoryInterfaceString)
	if err != nil {
		log.Fatal(err)
	}

	// 	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`
func (` + str.GetFirstCharacterOfString(moduleName) + ` ` + caser.String(moduleName) + `Db) ` + str.DashToCamel(methodName) + `() ([]` + moduleName + `Model.` + caser.String(moduleName) + `, error) {
	// Implement your retrieval logic here
	var ` + moduleName + ` []` + moduleName + `Model.` + caser.String(moduleName) + `
	` + moduleName + ` = append(` + moduleName + `, ` + moduleName + `Model.` + caser.String(moduleName) + `{Id: 1})
	return ` + moduleName + `, nil
}`
	infraestructureResult, err := find.FindFile("app/" + moduleName + "/infraestructure/")
	if err != nil {
		log.Fatal(err)
	}
	append.AppendDataToFile("app/"+moduleName+"/infraestructure/"+*infraestructureResult, infraestructureString)

	// Add endpoint to router.go
	input, err := os.ReadFile("router/router.go")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "//"+moduleName) {
			lines[i] = `	//` + moduleName + ` 
	router.GET("/` + endpointName + `/` + strings.ToLower(methodName) + `", ` + moduleName + `Controller.` + str.DashToCamel(methodName) + `)`
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
