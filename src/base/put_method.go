package base

import (
	"log"
	"os"
	"strings"

	utils "github.com/illud/gohex/src/utils/append"
	find "github.com/illud/gohex/src/utils/find"
	str "github.com/illud/gohex/src/utils/strings"
)

type PutMethodData struct {
	ModuleName      string
	StructName      string
	MethodNameLower string
	MethodFuncName  string
	EndpointName    string
	FirstChar       string
	EndpointCamelId string
}

func PutMethod(moduleName string, methodName string) {
	trackerResult := utils.ReadTrackerFile()
	var endpointName string

	for _, module := range trackerResult.Modules {
		if module.ModuleName == moduleName {
			endpointName = module.EndpointName
			break
		}
	}

	data := PutMethodData{
		ModuleName:      moduleName,
		StructName:      caser.String(moduleName),
		MethodNameLower: strings.ToLower(methodName),
		MethodFuncName:  str.DashToCamel(methodName),
		EndpointName:    endpointName,
		FirstChar:       str.GetFirstCharacterOfString(moduleName),
		EndpointCamelId: str.FormatHyphenToCamelCase(endpointName) + "Id",
	}

	// Controller
	controllerResult, err := find.FindFile("app/" + moduleName + "/aplication/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/controller.put.go.tmpl",
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
		"templates/endpoint/service.put.go.tmpl",
		"app/"+moduleName+"/domain/services/"+*serviceResult,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	repositoryResult, err := find.FindFile("app/" + moduleName + "/domain/repositories/")
	if err != nil {
		log.Fatal(err)
	}
	err = ReplaceLastCharacterWithTemplate(
		"app/"+moduleName+"/domain/repositories/"+*repositoryResult,
		"templates/endpoint/repository.put.go.tmpl",
		"}",
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Infra
	infraResult, err := find.FindFile("app/" + moduleName + "/infraestructure/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/infraestructure.put.go.tmpl",
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
			lines[i] = "\t//" + moduleName + "\n\trouter.PUT(\"/" + data.EndpointName + "/" + data.MethodNameLower + "/:" + data.EndpointCamelId + "\", " + moduleName + "Controller." + data.MethodFuncName + ")"
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
