package base

import (
	"log"
	"os"
	"strings"

	append "github.com/illud/gohex/src/utils/append"
	find "github.com/illud/gohex/src/utils/find"
	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func PutMethod(moduleName string, methodName string) {

	//Add data to controller.go
	controllerString :=
		`
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
// @Router /` + regex.StringToHyphen(moduleName) + `/` + methodName + `/{` + moduleName + `Id} [Put]
func ` + strings.Title(methodName) + `(c *gin.Context) {
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

	err = service.` + strings.Title(methodName) + `(` + moduleName + `Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "` + moduleName + ` updated",
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
func (s *Service) ` + strings.Title(methodName) + `(` + moduleName + `Id int) error {
	err := s.` + moduleName + `Repository.` + strings.Title(methodName) + `(` + moduleName + `Id)
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
		`	` + strings.Title(methodName) + `(` + moduleName + `Id int) error
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
func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) ` + strings.Title(methodName) + `(` + moduleName + `Id int) error {
	// Implement your update logic here
	return nil
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
	router.PUT("/` + moduleName + `/` + methodName + `/` + moduleName + `Id", ` + moduleName + `Controller.` + strings.Title(methodName) + `)`
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
