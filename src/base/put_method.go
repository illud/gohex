package base

import (
	"fmt"
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
// @Router /` + regex.StringToHyphen(moduleName) + `/{` + moduleName + `Id} [Put]
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

	result, err := service.` + strings.Title(methodName) + `(` + moduleName + `Id)
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
		fmt.Println("Error:", err)
	}
	// Write the data to the end of the file
	append.AppendDataToFile("app/"+moduleName+"/aplication/"+*controllerResult, controllerString)

	// 	//Add data to service.go
	servicesString :=
		`
func (s *Service) ` + strings.Title(methodName) + `(` + moduleName + `Id int) (*string, error) {
	result, err := s.` + moduleName + `Repository.` + strings.Title(methodName) + `(` + moduleName + `Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
`
	serviceResult, err := find.FindFile("app/" + moduleName + "/domain/services/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	append.AppendDataToFile("app/"+moduleName+"/domain/services/"+*serviceResult, servicesString)

	// 	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`	` + strings.Title(methodName) + `(` + moduleName + `Id int) (*string, error)
}`

	repositoryResult, err := find.FindFile("app/" + moduleName + "/domain/repositories/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = append.ReplaceLastCharacter("app/"+moduleName+"/domain/repositories/"+*repositoryResult, "}", repositoryInterfaceString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`
func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) ` + strings.Title(methodName) + `(` + moduleName + `Id int) (*string, error) {
	// Implement your update logic here
	var result = "` + strings.Title(methodName) + ` updated"
	return &result, nil
}
`
	infraestructureResult, err := find.FindFile("app/" + moduleName + "/infraestructure/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	append.AppendDataToFile("app/"+moduleName+"/infraestructure/"+*infraestructureResult, infraestructureString)
}
