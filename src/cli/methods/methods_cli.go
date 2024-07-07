package methods

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	base "github.com/illud/gohex/src/base"
	input "github.com/illud/gohex/src/cli/input"
)

var methods = []string{"POST", "GET", "PUT", "DELETE"}

type methodsModel struct {
	cursor int
	choice string
}

func (m methodsModel) Init() tea.Cmd {
	return nil
}

func (m methodsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = methods[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(methods) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(methods) - 1
			}
		}
	}

	return m, nil
}

func (m methodsModel) View() string {
	s := strings.Builder{}
	s.WriteString("please select your method\n\n")

	for i := 0; i < len(methods); i++ {
		if m.cursor == i {
			s.WriteString("▶  ")
		} else {
			s.WriteString("  ")
		}
		s.WriteString(methods[i])
		s.WriteString("\n")
	}

	return s.String()
}

// for new method
func Command(module string) {
	p := tea.NewProgram(methodsModel{})

	// Run returns the model as a tea.Model.
	m, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(methodsModel); ok && m.choice != "" {
		if m.choice == "POST" {
			fmt.Printf("\n")
			fmt.Println("Enter name for method: ")
			methodName := input.Input()
			base.PostMethod(module, methodName)
		}
		if m.choice == "GET" {
			fmt.Printf("\n")
			fmt.Println("Enter name for method: ")
			methodName := input.Input()
			base.GetMethod(module, methodName)
		}
		if m.choice == "PUT" {
			fmt.Printf("\n")
			fmt.Println("Enter name for method: ")
			methodName := input.Input()
			base.PutMethod(module, methodName)
		}
		if m.choice == "DELETE" {
			fmt.Printf("\n")
			fmt.Println("Enter name for method: ")
			methodName := input.Input()
			base.DeleteMethod(module, methodName)
		}
	}
}
