package methods

import (
	"fmt"
	"strings"

	"github.com/illud/gohex/src/utils"
	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func GetMethod(moduleName string, methodName string) {

	//Add data to controller.go
	controllerString :=
		`
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
func ` + strings.Title(methodName) + `(c *gin.Context) {
	result, err := service.` + strings.Title(methodName) + `()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
`
	// Write the data to the end of the file
	utils.AppendDataToFile("app/"+moduleName+"/aplication/"+moduleName+".controller.go", controllerString)

	// 	//Add data to service.go
	servicesString :=
		`
func (s *Service) ` + strings.Title(methodName) + `() ([]*` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	result, err := s.` + moduleName + `Repository.` + strings.Title(methodName) + `()
	if err != nil {
		return nil, err
	}
	return result, nil
}`
	utils.AppendDataToFile("app/"+moduleName+"/domain/services/"+moduleName+".service.go", servicesString)

	// 	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`	` + strings.Title(methodName) + `() ([]*` + moduleName + `Model.` + strings.Title(moduleName) + `, error)
}`

	err := utils.ReplaceLastCharacter("app/"+moduleName+"/domain/repositories/"+moduleName+".repository.go", "}", repositoryInterfaceString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`
func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) ` + strings.Title(methodName) + `() ([]*` + moduleName + `Model.` + strings.Title(moduleName) + `, error) {
	// Implement your retrieval logic here
	var ` + moduleName + ` []*` + moduleName + `Model.` + strings.Title(moduleName) + `
	` + moduleName + ` = append(` + moduleName + `, &` + moduleName + `Model.` + strings.Title(moduleName) + `{Id: 1})
	return ` + moduleName + `, nil
}`
	utils.AppendDataToFile("app/"+moduleName+"/infraestructure/"+moduleName+".db.go", infraestructureString)
}
