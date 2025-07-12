package base

import "log"

func BaseDbClient(clientName string) {
	templatePath := ""
	switch clientName {
	case "mysql":
		templatePath = "templates/database/mysql.tmpl"
	case "gorm":
		templatePath = "templates/database/gorm.tmpl"
	default:
		log.Fatalf("Unsupported client name: %s", clientName)
	}

	// This function is used to create a database client based on the provided client name.
	writeTemplateToFile(templatePath, "adapters/database/db.go", struct{}{})

	AppendDbConnectionToMain()
}
