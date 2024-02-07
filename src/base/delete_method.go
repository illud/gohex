package base

import (
	"fmt"
	"strings"

	append "github.com/illud/gohex/src/utils/append"
	find "github.com/illud/gohex/src/utils/find"
	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func DeleteMethod(moduleName string, methodName string) {

	//Add data to controller.go
	controllerString :=
		`
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
func ` + strings.Title(methodName) + `(c *gin.Context) {
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
	// Implement your deletion logic here
	var result = "` + strings.Title(methodName) + ` deleted"
	return &result, nil
}
`
	infraestructureResult, err := find.FindFile("app/" + moduleName + "/infraestructure/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	append.AppendDataToFile("app/"+moduleName+"/infraestructure/"+*infraestructureResult, infraestructureString)
}
