package base

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	utils "github.com/illud/gohex/src/utils/append"
	str "github.com/illud/gohex/src/utils/strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.Und)

// ADD controller to router.go crud
func AppendToRoutingCrud(moduleName string, moduleNotModify string) {
	// add data to tracker
	err := utils.AddDataToTrackerFile(moduleName, str.ToKebabCase(moduleNotModify))
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("router/router.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	` + moduleName + `Controller "github.com/` + currentDirName + `/app/` + moduleName + `/aplication"`
		}

		if strings.Contains(line, "return router") {
			lines[i] = ` //` + moduleName + `
	router.POST("/` + str.ToKebabCase(moduleNotModify) + `", ` + moduleName + `Controller.Create` + caser.String(moduleName) + `)
	router.GET("/` + str.ToKebabCase(moduleNotModify) + `", ` + moduleName + `Controller.Get` + caser.String(moduleName) + `)
	router.GET("/` + str.ToKebabCase(moduleNotModify) + `/:` + str.FormatSnakeCaseToCamelCase(moduleNotModify) + `Id", ` + moduleName + `Controller.GetOne` + caser.String(moduleName) + `)
	router.PUT("/` + str.ToKebabCase(moduleNotModify) + `/:` + str.FormatSnakeCaseToCamelCase(moduleNotModify) + `Id", ` + moduleName + `Controller.Update` + caser.String(moduleName) + `)
	router.DELETE("/` + str.ToKebabCase(moduleNotModify) + `/:` + str.FormatSnakeCaseToCamelCase(moduleNotModify) + `Id", ` + moduleName + `Controller.Delete` + caser.String(moduleName) + `)

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("router/router.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt router/router.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt router/router.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}

// ADD controller to router.go simple
func AppendToRoutingSimple(moduleName string, moduleNotModify string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("routing/routing.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	` + moduleName + `Controller "github.com/` + currentDirName + `/controller/` + moduleName + `"`
		}

		if strings.Contains(line, "return router") {
			lines[i] = ` //` + moduleName + `
	router.GET("/` + str.ToKebabCase(moduleNotModify) + `", ` + moduleName + `Controller.Get` + caser.String(moduleName) + `)

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("routing/routing.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt routing/routing.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt routing/routing.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}

// ADD db conection to main.go
func AppendDbConnectionToMain() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	input, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "import (") || strings.Contains(line, "import(") {
			lines[i] = `import (
	db "` + currentDirName + `/adapters/database"`
		}

		if strings.Contains(line, "router.Router().Run") {
			lines[i] = ` //Connect to database
			db.Connect()

` + lines[i] + ``
		}

	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile("main.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//format main.go
	if runtime.GOOS == "windows" {
		installDependencies := exec.Command("cmd", "/c", "go fmt main.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}

	if runtime.GOOS == "linux" {
		installDependencies := exec.Command("sh", "/c", "go fmt main.go")

		_, err = installDependencies.Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		// fmt.Println(string(installDependenciesOut))
	}
}
