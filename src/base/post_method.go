package base

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	append "github.com/illud/gohex/src/utils/append"
	utils "github.com/illud/gohex/src/utils/append"
	find "github.com/illud/gohex/src/utils/find"
	str "github.com/illud/gohex/src/utils/strings"
)

func AppendTemplateToFile(templatePath string, outputPath string, data interface{}) error {
	// ✅ 1. Lee el contenido directamente desde el FS embebido
	tmplContent, err := TmplFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading embedded template %s: %w", templatePath, err)
	}

	// ✅ 2. Parsea el contenido
	tmpl, err := template.New("tmpl").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("error parsing template content %s: %w", templatePath, err)
	}

	// ✅ 3. Ejecuta la plantilla
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, data); err != nil {
		return fmt.Errorf("error rendering template %s: %w", templatePath, err)
	}

	return append.AppendDataToFile(outputPath, rendered.String())
}

func ReplaceLastCharacterWithTemplate(filePath, templatePath, oldChar string, data interface{}) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	contentStr := string(content)

	lastIndex := strings.LastIndex(contentStr, oldChar)
	if lastIndex == -1 {
		return fmt.Errorf("'%s' not found in the file", oldChar)
	}

	// Lee y parsea desde el embed
	tmplContent, err := TmplFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading embedded template %s: %w", templatePath, err)
	}
	tmpl, err := template.New("tmpl").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("error parsing template content %s: %w", templatePath, err)
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, data); err != nil {
		return fmt.Errorf("error rendering template %s: %w", templatePath, err)
	}

	newContent := contentStr[:lastIndex] + rendered.String() + contentStr[lastIndex+1:]

	return os.WriteFile(filePath, []byte(newContent), 0644)
}

type PostMethodData struct {
	ModuleName      string
	StructName      string
	MethodNameLower string
	MethodFuncName  string
	EndpointName    string
	FirstChar       string
}

func PostMethod(moduleName string, methodName string) {
	trackerResult := utils.ReadTrackerFile()
	var endpointName string

	for _, module := range trackerResult.Modules {
		if module.ModuleName == moduleName {
			endpointName = module.EndpointName
			break
		}
	}

	data := PostMethodData{
		ModuleName:      moduleName,
		StructName:      caser.String(moduleName),
		MethodNameLower: strings.ToLower(methodName),
		MethodFuncName:  str.DashToCamel(methodName),
		EndpointName:    endpointName,
		FirstChar:       str.GetFirstCharacterOfString(moduleName),
	}

	// Controller
	controllerResult, err := find.FindFile("app/" + moduleName + "/aplication/")
	if err != nil {
		log.Fatal(err)
	}
	err = AppendTemplateToFile(
		"templates/endpoint/controller.post.go.tmpl",
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
		"templates/endpoint/service.post.go.tmpl",
		"app/"+moduleName+"/domain/services/"+*serviceResult,
		data,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Repository Interface (reemplazar la llave final)
	repositoryResult, err := find.FindFile("app/" + moduleName + "/domain/repositories/")
	if err != nil {
		log.Fatal(err)
	}
	err = ReplaceLastCharacterWithTemplate(
		"app/"+moduleName+"/domain/repositories/"+*repositoryResult,
		"templates/endpoint/repository.post.go.tmpl",
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
		"templates/endpoint/infraestructure.post.go.tmpl",
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
			lines[i] = "\t//" + moduleName + "\n\trouter.POST(\"/" + data.EndpointName + "/" + data.MethodNameLower + "\", " + moduleName + "Controller." + data.MethodFuncName + ")"

		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
