package base

import (
	"log"
	"os"
	"strings"

	utils "github.com/illud/gohex/src/utils/append"
	"github.com/illud/gohex/src/utils/find"
	str "github.com/illud/gohex/src/utils/strings"
)

type DeleteMethodData struct {
	ModuleName      string
	StructName      string
	MethodNameLower string
	MethodFuncName  string
	EndpointName    string
	FirstChar       string
	ParamId         string
}

func DeleteMethod(moduleName string, methodName string) {
	trackerResult := utils.ReadTrackerFile()
	var endpointName string

	for _, module := range trackerResult.Modules {
		if module.ModuleName == moduleName {
			endpointName = module.EndpointName
			break
		}
	}

	data := DeleteMethodData{
		ModuleName:      moduleName,
		StructName:      caser.String(moduleName),
		MethodNameLower: strings.ToLower(methodName),
		MethodFuncName:  str.DashToCamel(methodName),
		EndpointName:    endpointName,
		FirstChar:       str.GetFirstCharacterOfString(moduleName),
		ParamId:         str.FormatHyphenToCamelCase(endpointName) + "Id",
	}

	// Controller
	controllerResult, err := find.FindFile("app/" + moduleName + "/aplication/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/controller.delete.go.tmpl",
		"app/"+moduleName+"/aplication/"+*controllerResult,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Service
	serviceResult, err := find.FindFile("app/" + moduleName + "/domain/services/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/service.delete.go.tmpl",
		"app/"+moduleName+"/domain/services/"+*serviceResult,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Repository interface
	repositoryResult, err := find.FindFile("app/" + moduleName + "/domain/repositories/")
	if err != nil {
		log.Fatal(err)
	}
	err = ReplaceLastCharacterWithTemplate(
		"app/"+moduleName+"/domain/repositories/"+*repositoryResult,
		"templates/endpoint/repository.delete.go.tmpl",
		"}",
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Infraestructure
	infraResult, err := find.FindFile("app/" + moduleName + "/infraestructure/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/infraestructure.delete.go.tmpl",
		"app/"+moduleName+"/infraestructure/"+*infraResult,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Router
	input, err := os.ReadFile("router/router.go")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "//"+moduleName) {
			lines[i] = "\t//" + moduleName + "\n\trouter.DELETE(\"/" + data.EndpointName + "/" + data.MethodNameLower + "/:" + data.ParamId + "\", " + moduleName + "Controller." + data.MethodFuncName + ")"
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
