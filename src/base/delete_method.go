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

func DeleteMethod(moduleName string, methodName string) {
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
// Delete ` + strings.Title(moduleName) + `
// @Summary Delete ` + strings.Title(moduleName) + `
// @Schemes
// @Description Delete ` + strings.Title(moduleName) + `
// @Tags ` + strings.Title(moduleName) + `
// @Security BearerAuth
// @Param ` + str.FormatHyphenToCamelCase(endpointName) + `Id path int true "` + str.FormatHyphenToCamelCase(endpointName) + `Id"
// @Accept json
// @Produce json
// @Success 200
// @Router /` + endpointName + `/` + strings.ToLower(methodName) + `/{` + str.FormatHyphenToCamelCase(endpointName) + `Id} [Delete]
func ` + strings.Title(str.DashToCamel(methodName)) + `(c *gin.Context) {
	` + str.FormatHyphenToCamelCase(endpointName) + `Id, err := strconv.Atoi(c.Param("` + str.FormatHyphenToCamelCase(endpointName) + `Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.` + strings.Title(str.DashToCamel(methodName)) + `(` + str.FormatHyphenToCamelCase(endpointName) + `Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "` + moduleName + ` deleted",
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
func (s *Service) ` + strings.Title(str.DashToCamel(methodName)) + `(` + str.FormatHyphenToCamelCase(endpointName) + `Id int) error {
	err := s.` + moduleName + `Repository.` + strings.Title(str.DashToCamel(methodName)) + `(` + str.FormatHyphenToCamelCase(endpointName) + `Id)
	if err != nil {
		return err
	}
	return nil
}
`
	serviceResult, err := find.FindFile("app/" + moduleName + "/domain/services/")
	if err != nil {
		log.Fatal(err)
	}
	append.AppendDataToFile("app/"+moduleName+"/domain/services/"+*serviceResult, servicesString)

	// 	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`	` + strings.Title(str.DashToCamel(methodName)) + `(` + str.FormatHyphenToCamelCase(endpointName) + `Id int) error
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
func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) ` + strings.Title(str.DashToCamel(methodName)) + `(` + str.FormatHyphenToCamelCase(endpointName) + `Id int) error {
	// Implement your deletion logic here
	return  nil
}
`
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
	router.DELETE("/` + endpointName + `/` + strings.ToLower(methodName) + `/:` + str.FormatHyphenToCamelCase(endpointName) + `Id", ` + moduleName + `Controller.` + strings.Title(str.DashToCamel(methodName)) + `)`
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
