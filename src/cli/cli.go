package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/illud/gohex/src/base"
	input "github.com/illud/gohex/src/cli/input"
	"github.com/schollz/progressbar/v3"
)

var choices = []string{"New Project", "Module", "DB Service"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Choose a option\n\n")
	s.WriteString("⇡/⇣ to select\n\n")
	s.WriteString("please use snake_case when the module name consist of two or more words\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("▶  ")
		} else {
			s.WriteString("  ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func Command() {

	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		// fmt.Printf("\n---\nYou chose %s!\n", m.choice)

		if m.choice == "New Project" {

			fmt.Printf("\n")
			fmt.Println("Enter Project Name: ")
			folder := input.Input()

			folderName := strings.ToLower(folder)

			fmt.Printf("\n")
			//Project
			bar := progressbar.Default(41)

			os.MkdirAll(folderName, os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			os.Create(folderName + "/main.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Router
			os.MkdirAll(folderName+"/router", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			os.Create(folderName + "/router/router.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//.env
			os.Create(folderName + "/.env")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//.gitignore
			os.Create(folderName + "/.gitignore")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//README
			os.Create(folderName + "/README")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Env
			os.MkdirAll(folderName+"/env", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			os.Create(folderName + "/env/env.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App folder
			os.MkdirAll(folderName+"/app", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//app/tasks
			os.MkdirAll(folderName+"/app/tasks", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//app/tasks/aplication
			os.MkdirAll(folderName+"/app/tasks/aplication", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//app/aplication/tasks/tasks.controller.go
			os.Create(folderName + "/app/tasks/aplication/tasks.controller.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App tasks domain folder
			os.MkdirAll(folderName+"/app/tasks/domain", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App tasks domain/models folder
			os.MkdirAll(folderName+"/app/tasks/domain/models", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//model
			os.Create(folderName + "/app/tasks/domain/models/tasks.model.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App tasks domain/repositories folder
			os.MkdirAll(folderName+"/app/tasks/domain/repositories", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//repository
			os.Create(folderName + "/app/tasks/domain/repositories/tasks.repository.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App tasks domain/services folder
			os.MkdirAll(folderName+"/app/tasks/domain/services", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//services
			os.Create(folderName + "/app/tasks/domain/services/tasks.service.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Infraestructure
			os.MkdirAll(folderName+"/app/tasks/infraestructure", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//databases/tasks.go
			os.Create(folderName + "/app/tasks/infraestructure/tasks.db.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Data
			os.MkdirAll(folderName+"/data", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//data/db
			os.Create(folderName + "/data/db.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//helpers
			os.MkdirAll(folderName+"/helpers", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//helpers/async.go
			os.Create(folderName + "/helpers/async.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//helpers/errors.go
			os.Create(folderName + "/helpers/errors.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//jswt
			os.MkdirAll(folderName+"/jswt", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//jswt/jswt.go
			os.Create(folderName + "/jswt/jswt.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//bcrypt
			os.MkdirAll(folderName+"/bcrypt", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//bcrypt/bcrypt.go
			os.Create(folderName + "/bcrypt/bcrypt.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//E2E TEST FOLDER
			os.MkdirAll(folderName+"/e2e", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//E2E TEST TASKS FOLDER
			os.MkdirAll(folderName+"/e2e/tasks", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			os.Create(folderName + "/e2e/tasks/gettasks_test.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Create base files data
			base.BaseData(folderName)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			if runtime.GOOS == "windows" {
				cmd := exec.Command("cmd", "/c", "go mod init "+folderName)
				cmd.Dir = folderName

				//INSTALL DEPENDENCIES
				_, err = cmd.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(out))

				//Install swago
				// fmt.Println("	")
				// fmt.Println("executing go install github.com/swaggo/swag/cmd/swag@latest")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installSwagDependencies := exec.Command("cmd", "/c", "go install github.com/swaggo/swag/cmd/swag@latest")

				//INSTALL DEPENDENCIES
				_, err = installSwagDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				//Install fresh
				// fmt.Println("	")
				// fmt.Println("executing go install github.com/pilu/fresh@latest")
				// fmt.Println("	")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installFreshDependencies := exec.Command("cmd", "/c", "go install github.com/pilu/fresh@latest")

				//INSTALL DEPENDENCIES
				_, err = installFreshDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				//Install gotestsum
				// fmt.Println("	")
				// fmt.Println("executing go install gotest.tools/gotestsum@latest")
				// fmt.Println("	")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installGotestsumDependencies := exec.Command("cmd", "/c", "go install gotest.tools/gotestsum@latest")

				//INSTALL gotestsum DEPENDENCIES
				_, err = installGotestsumDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				installDependencies := exec.Command("cmd", "/c", "go get -d ./...")
				installDependencies.Dir = folderName

				//SWAG INIT
				swagInit := exec.Command("cmd", "/c", "swag init")
				swagInit.Dir = folderName

				_, err = swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(swagInitOut))

				//INSTALL DEPENDENCIES
				_, err = installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installDependenciesOut))

				//INSTALL TEST DEPENDENCIES
				installTestDependencies := exec.Command("cmd", "/c", "go get -t ./...")
				installTestDependencies.Dir = folderName

				_, err = installTestDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installTestDependenciesOut))

			}

			if runtime.GOOS == "linux" {
				cmd := exec.Command("sh", "/c", "go mod init github.com/"+folderName)
				cmd.Dir = folderName

				//INSTALL DEPENDENCIES
				_, err = cmd.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(out))

				//Install swago
				// fmt.Println("")
				// fmt.Println("executing go install github.com/swaggo/swag/cmd/swag@latest")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installSwagDependencies := exec.Command("sh", "/c", "go install github.com/swaggo/swag/cmd/swag@latest")

				//INSTALL DEPENDENCIES
				_, err = installSwagDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				//Install fresh
				// fmt.Println("	")
				// fmt.Println("executing go install github.com/pilu/fresh@latest")
				// fmt.Println("	")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installFreshDependencies := exec.Command("sh", "/c", "go install github.com/pilu/fresh@latest")

				//INSTALL DEPENDENCIES
				_, err = installFreshDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				//Install gotestsum
				// fmt.Println("	")
				// fmt.Println("executing go install gotest.tools/gotestsum@latest")
				// fmt.Println("	")
				bar.Add(1)
				time.Sleep(40 * time.Millisecond)

				installgotestsumDependencies := exec.Command("sh", "/c", "go install gotest.tools/gotestsum@latest")

				//INSTALL gotestsum DEPENDENCIES
				_, err = installgotestsumDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}

				installDependencies := exec.Command("sh", "/c", "go get -d ./...")
				installDependencies.Dir = folderName

				//SWAG INIT
				swagInit := exec.Command("sh", "/c", "swag init")
				swagInit.Dir = folderName

				_, err = swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(swagInitOut))

				//INSTALL DEPENDENCIES
				_, err = installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installDependenciesOut))

				//INSTALL TEST DEPENDENCIES
				installTestDependencies := exec.Command("sh", "/c", "go get -t ./...")
				installTestDependencies.Dir = folderName

				_, err = installTestDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installTestDependenciesOut))
			}

			bar.Add(3)
			time.Sleep(40 * time.Millisecond)

			fmt.Println("")

			//Display usage
			fmt.Println(" | get started ")
			fmt.Println(" | cd ", folderName)
			fmt.Println(" | go run main.go ")
			fmt.Println("")

		}

		if m.choice == "Module" {
			fmt.Printf("\n")
			fmt.Println("Enter Module Name: ")
			module := input.Input()

			moduleNameNoSnakeCase := strings.Replace(module, "_", "", -1)
			moduleName := strings.ToLower(moduleNameNoSnakeCase)
			moduleNameSnakeCase := strings.ToLower(module)

			bar := progressbar.Default(13)

			//app
			os.MkdirAll("app/"+moduleName+"", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//plication
			os.MkdirAll("app/"+moduleName+"/aplication", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//module.controller.go
			os.Create("app/" + moduleName + "/aplication/" + moduleNameSnakeCase + ".controller.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App module domain folder
			os.MkdirAll("app/"+moduleName+"/domain", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App module domain/models folder
			os.MkdirAll("app/"+moduleName+"/domain/models", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//model
			os.Create("app/" + moduleName + "/domain/models/" + moduleNameSnakeCase + ".model.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App module domain/repositories folder
			os.MkdirAll("app/"+moduleName+"/domain/repositories", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//repository
			os.Create("app/" + moduleName + "/domain/repositories/" + moduleNameSnakeCase + ".repository.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//App module domain/services folder
			os.MkdirAll("app/"+moduleName+"/domain/services", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//services
			os.Create("app/" + moduleName + "/domain/services/" + moduleNameSnakeCase + ".service.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Infraestructure
			os.MkdirAll("app/"+moduleName+"/infraestructure", os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//infraestructure/module.go
			os.Create("app/" + moduleName + "/infraestructure/" + moduleNameSnakeCase + ".db.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//E2E TEST FOLDER
			os.MkdirAll("e2e/"+moduleName, os.ModePerm)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			os.Create("e2e/" + moduleName + "/get_" + moduleNameSnakeCase + "_test.go")
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Generates module with crud data
			base.BaseModuleCrud(moduleName, moduleNameSnakeCase)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//Append controller to routing.go file
			base.AppendToRoutingCrud(moduleName)
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

			//SWAG INIT Windows
			if runtime.GOOS == "windows" {
				swagInit := exec.Command("cmd", "/c", "swag init")

				_, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(swagInitOut))
			}

			//SWAG INIT Linux
			if runtime.GOOS == "linux" {
				swagInit := exec.Command("sh", "/c", "swag init")

				_, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(swagInitOut))
			}

			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
		}

		if m.choice == "DB Service" {
			fmt.Printf("\n")
			fmt.Println("Enter DB(mysql, gorm) Name: ")
			flagName := input.Input()

			bar := progressbar.Default(1)

			if flagName == "mysql" {
				base.BaseDbClient("mysql")
			}
			if flagName == "gorm" {
				base.BaseDbClient("gorm")
			}
			// if flagName == "prisma" {
			// 	base.BaseDbClient("prisma")
			// 	//create/schema.prisma
			// 	os.MkdirAll("data/prisma/db", os.ModePerm)

			// 	os.Create("data/prisma/schema.prisma")

			// 	fmt.Println("")

			// 	fmt.Println("To get this up and running in your database, we use the Prisma migration tool migrate to create and migrate our database:")
			// 	fmt.Println("sync the database with your schema go run github.com/prisma/prisma-client-go migrate dev --name init")
			// 	fmt.Println("After the migration, the Prisma Client Go client is automatically generated in your project. If you just want to re-generate the client, run go run github.com/prisma/prisma-client-go generate.")

			// 	fmt.Println("For more visit https://github.com/prisma/prisma-client-go")

			// 	//Install db DEPENDENCIES
			// 	if runtime.GOOS == "windows" {
			// 		fmt.Println("")
			// 		fmt.Println("executing go get github.com/prisma/prisma-client-go")

			// 		installDependencies := exec.Command("cmd", "/c", "go get github.com/prisma/prisma-client-go")

			// 		//INSTALL DEPENDENCIES
			// 		_, err = installDependencies.Output()
			// 		if err != nil {
			// 			os.Stderr.WriteString(err.Error())
			// 		}
			// 		// fmt.Println(string(installDependenciesOut))

			// 		//Run prisma init
			// 		installPrismaDependencies := exec.Command("cmd", "/c", "go run github.com/prisma/prisma-client-go migrate dev --name init")

			// 		//INSTALL DEPENDENCIES
			// 		_, err = installPrismaDependencies.Output()
			// 		if err != nil {
			// 			os.Stderr.WriteString(err.Error())
			// 		}
			// 		// fmt.Println(string(installPrismaDependenciesOut))
			// 	}

			// 	if runtime.GOOS == "linux" {
			// 		fmt.Println("")
			// 		fmt.Println("executing go get github.com/prisma/prisma-client-go")

			// 		installDependencies := exec.Command("sh", "/c", "go get github.com/prisma/prisma-client-go.")

			// 		//INSTALL DEPENDENCIES
			// 		_, err = installDependencies.Output()
			// 		if err != nil {
			// 			os.Stderr.WriteString(err.Error())
			// 		}
			// 		// fmt.Println(string(installDependenciesOut))

			// 		//Run prisma init
			// 		installPrismaDependencies := exec.Command("sh", "/c", "go run github.com/prisma/prisma-client-go migrate dev --name init")

			// 		//INSTALL DEPENDENCIES
			// 		_, err = installPrismaDependencies.Output()
			// 		if err != nil {
			// 			os.Stderr.WriteString(err.Error())
			// 		}
			// 		// fmt.Println(string(installPrismaDependenciesOut))
			// 	}

			// }

			//Install db DEPENDENCIES
			if runtime.GOOS == "windows" {
				installDependencies := exec.Command("cmd", "/c", "go get -d ./...")

				//INSTALL DEPENDENCIES
				_, err = installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installDependenciesOut))
			}

			if runtime.GOOS == "linux" {
				installDependencies := exec.Command("sh", "/c", "go get -d ./...")

				//INSTALL DEPENDENCIES
				_, err = installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				// fmt.Println(string(installDependenciesOut))
			}

			bar.Add(1)
		}

		if m.choice == "Documentation" {
			path, _ := filepath.Abs("README.md")
			source, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}

			result := markdown.Render(string(source), 80, 6)

			fmt.Println(string(result))
		}
	}
}
