package base

import (
	"bytes"
	"html/template"
	"os"
)

func writeTemplateToFile(embedPath, outputPath string, data any) {
	// Read the template content from the embedded filesystem
	tmplContent, err := TmplFS.ReadFile(embedPath)
	if err != nil {
		panic("error reading embedded template: " + err.Error())
	}

	// Create a new template and parse the content
	tmpl := template.Must(template.New("embedded").Parse(string(tmplContent)))

	// Execute the template with the provided data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("error executing template: " + err.Error())
	}

	// Ensure the directory exists
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		panic("error writing output file: " + err.Error())
	}
}

func BaseData(folderName string) {
	//Add data to tracker.json
	writeTemplateToFile("templates/base/tracker.json.tmpl", folderName+"/tracker.json", struct {
		ModuleName   string
		EndpointName string
	}{
		ModuleName:   "tasks",
		EndpointName: "tasks",
	})

	//Add data to main.go
	writeTemplateToFile("templates/base/main.go.tmpl", folderName+"/main.go", struct {
		FolderName string
	}{
		FolderName: folderName,
	})

	//Add data to router.go
	writeTemplateToFile("templates/base/router.go.tmpl", folderName+"/router/router.go", struct {
		FolderName string
	}{
		FolderName: folderName,
	})

	//Add data to .env
	writeTemplateToFile("templates/base/.env.tmpl", folderName+"/.env", struct {
		Port    int
		Version string
	}{
		Port:    5000,
		Version: "1.0.0",
	})

	//Add data to .gitignore
	writeTemplateToFile("templates/base/.gitignore.tmpl", folderName+"/.gitignore", nil)

	//Add data to README
	writeTemplateToFile("templates/base/README.tmpl", folderName+"/README.md", nil)

	//Add data to env.go
	writeTemplateToFile("templates/base/env.go.tmpl", folderName+"/env/env.go", nil)

	//Add data to task-controller.go
	writeTemplateToFile(
		"templates/base/tasks.controller.go.tmpl",
		folderName+"/app/tasks/aplication/tasks.controller.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

	//Add data to models/task.model.go
	writeTemplateToFile(
		"templates/base/tasks.model.go.tmpl",
		folderName+"/app/tasks/domain/models/tasks.model.go",
		nil,
	)

	//Add data to repositories/tasks.repository.go
	writeTemplateToFile(
		"templates/base/tasks.repository.go.tmpl",
		folderName+"/app/tasks/domain/repositories/tasks.repository.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

	//Add data to task.service.go
	writeTemplateToFile(
		"templates/base/tasks.service.go.tmpl",
		folderName+"/app/tasks/domain/services/tasks.service.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

	//Add data to tasks.db.go
	writeTemplateToFile(
		"templates/base/tasks.db.go.tmpl",
		folderName+"/app/tasks/infraestructure/tasks.db.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

	//Add data to data/db.go
	writeTemplateToFile(
		"templates/base/db.go.tmpl",
		folderName+"/adapters/database/db.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

	//Add data to adapters/bcrypt/bcrypt.go
	writeTemplateToFile(
		"templates/base/bcrypt.go.tmpl",
		folderName+"/adapters/bcrypt/bcrypt.go",
		struct{}{},
	)

	// jwt
	writeTemplateToFile(
		"templates/base/jwt.go.tmpl",
		folderName+"/adapters/jwt/jwt.go",
		struct{}{},
	)

	// ERRORS
	writeTemplateToFile(
		"templates/base/errors.go.tmpl",
		folderName+"/helpers/errors.go",
		struct{}{},
	)

	// getTasks_test.go
	writeTemplateToFile(
		"templates/base/getTasks_test.go.tmpl",
		folderName+"/e2e/tasks/getTasks_test.go",
		struct {
			FolderName string
		}{
			FolderName: folderName,
		},
	)

}
