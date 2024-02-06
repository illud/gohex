package methods

import (
	"fmt"
	"strings"

	"github.com/illud/gohex/src/utils"
	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func PostMethod(moduleName string, methodName string) {

	//Add data to controller.go
	controllerString :=
		`
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
func ` + strings.Title(methodName) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Model.` + strings.Title(moduleName) + `
	if err := c.ShouldBindJSON(&` + moduleName + `); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
func (s *Service) ` + strings.Title(methodName) + `() (*string, error) {
	result, err := s.` + moduleName + `Repository.` + strings.Title(methodName) + `()
	if err != nil {
		return nil, err
	}
	return result, nil
}
`
	utils.AppendDataToFile("app/"+moduleName+"/domain/services/"+moduleName+".service.go", servicesString)

	// 	//Add data to module/infraestructure/module.db.go
	repositoryInterfaceString :=
		`	` + strings.Title(methodName) + `() (*string, error)
}`

	err := utils.ReplaceLastCharacter("app/"+moduleName+"/domain/repositories/"+moduleName+".repository.go", "}", repositoryInterfaceString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 	//Add data to module/infraestructure/module.db.go
	infraestructureString :=
		`
func (` + str.GetFirstCharacterOfString(moduleName) + ` *` + strings.Title(moduleName) + `Db) ` + strings.Title(methodName) + `() (*string, error) {
	// Implement your creation logic here
	var result = "` + strings.Title(moduleName) + ` created"
	return &result, nil
}
`
	utils.AppendDataToFile("app/"+moduleName+"/infraestructure/"+moduleName+".db.go", infraestructureString)
}
