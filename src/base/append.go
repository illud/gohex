package base

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	regex "github.com/illud/gohex/src/utils/regex"
)

// ADD controller to router.go crud
func AppendToRoutingCrud(moduleName string) {
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
	router.POST("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Create` + strings.Title(moduleName) + `)
	router.GET("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Get` + strings.Title(moduleName) + `)
	router.GET("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.GetOne` + strings.Title(moduleName) + `)
	router.PUT("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.Update` + strings.Title(moduleName) + `)
	router.DELETE("/` + regex.StringToHyphen(moduleName) + `/:` + moduleName + `Id", ` + moduleName + `Controller.Delete` + strings.Title(moduleName) + `)

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
func AppendToRoutingSimple(moduleName string) {
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
	router.GET("/` + regex.StringToHyphen(moduleName) + `", ` + moduleName + `Controller.Get` + strings.Title(moduleName) + `)

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
