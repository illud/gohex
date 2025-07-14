package base

import (
	"log"
	"os"
	"runtime"
	"strings"

	regex "github.com/illud/gohex/src/utils/regex"
	str "github.com/illud/gohex/src/utils/strings"
)

func BaseModuleCrud(moduleName string, moduleNameSnakeCase string, moduleNotModify string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	//Add data to controller.go
	writeTemplateToFile(
		"templates/module/controller.go.tmpl",
		"app/"+moduleName+"/aplication/"+moduleNameSnakeCase+".controller.go",
		struct {
			CurrentDirName  string
			ModuleName      string
			ModuleNameTitle string
			ModuleCamel     string
		}{
			CurrentDirName:  currentDirName,
			ModuleName:      moduleName,
			ModuleNameTitle: caser.String(moduleName),
			ModuleCamel:     str.FormatSnakeCaseToCamelCase(moduleName),
		},
	)

	//Add data to moduleName.modle.go
	writeTemplateToFile(
		"templates/module/model.go.tmpl",
		"app/"+moduleName+"/domain/models/"+moduleNameSnakeCase+".model.go",
		struct {
			StructName string
		}{
			StructName: caser.String(moduleName),
		},
	)

	//Add data to service.go
	writeTemplateToFile(
		"templates/module/service.go.tmpl",
		"app/"+moduleName+"/domain/services/"+moduleNameSnakeCase+".service.go",
		struct {
			ModuleName     string
			CurrentDirName string
			StructName     string
			ParamName      string
		}{
			ModuleName:     moduleName,
			CurrentDirName: currentDirName,
			StructName:     caser.String(moduleName),
			ParamName:      str.FormatSnakeCaseToCamelCase(moduleNotModify),
		},
	)

	writeTemplateToFile(
		"templates/module/repository.go.tmpl",
		"app/"+moduleName+"/domain/repositories/"+moduleNameSnakeCase+".repository.go",
		struct {
			ModuleName     string
			CurrentDirName string
			StructName     string
			IdName         string
		}{
			ModuleName:     moduleName,
			CurrentDirName: currentDirName,
			StructName:     caser.String(moduleName),
			IdName:         str.FormatSnakeCaseToCamelCase(moduleNotModify) + "Id",
		},
	)

	//Add data to module/infraestructure/module.db.go
	writeTemplateToFile(
		"templates/module/infraestructure.go.tmpl",
		"app/"+moduleName+"/infraestructure/"+moduleNameSnakeCase+".db.go",
		struct {
			ModuleName     string
			CurrentDirName string
			StructName     string
			ParamName      string
			FirstChar      string
		}{
			ModuleName:     moduleName,
			CurrentDirName: currentDirName,
			StructName:     caser.String(moduleName),
			ParamName:      str.FormatSnakeCaseToCamelCase(moduleNotModify),
			FirstChar:      str.GetFirstCharacterOfString(moduleName),
		},
	)

	// TEST
	writeTemplateToFile(
		"templates/module/test.go.tmpl",
		"e2e/"+moduleName+"/get_"+moduleNameSnakeCase+"_test.go",
		struct {
			ModuleName     string
			CurrentDirName string
			StructName     string
			KebabCase      string
		}{
			ModuleName:     moduleName,
			CurrentDirName: currentDirName,
			StructName:     caser.String(moduleName),
			KebabCase:      regex.StringToHyphen(moduleName),
		},
	)

}
