package endpoint

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/illud/gohex/src/cli/methods"
)

var endpointChoices = []string{"New Project", "Module", "Endpoint", "DB Service"}

type endpointModel struct {
	cursor int
	choice string
}

func (m endpointModel) Init() tea.Cmd {
	return nil
}

func (m endpointModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = endpointChoices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(endpointChoices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(endpointChoices) - 1
			}
		}
	}

	return m, nil
}

func (m endpointModel) View() string {
	s := strings.Builder{}
	s.WriteString("please select your module\n\n")
	s.WriteString("please use hyphen-case when the module name consist of two or more words\n\n")

	for i := 0; i < len(endpointChoices); i++ {
		if m.cursor == i {
			s.WriteString("▶  ")
		} else {
			s.WriteString("  ")
		}
		s.WriteString(endpointChoices[i])
		s.WriteString("\n")
	}

	return s.String()
}

// for new endpoint
func Command(folders []string) {
	endpointChoices = folders
	var module string
	p := tea.NewProgram(endpointModel{})

	// Run returns the model as a tea.Model.
	m, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(endpointModel); ok && m.choice != "" {
		if m.choice == m.choice {
			module = m.choice
			methods.Command(module) // for new method
		}
	}
}
