package endpoint

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/illud/gohex/src/cli/methods"
)

type endpointModel struct {
	cursor  int
	choices []string
	choice  string
}

// Create a new instance of endpointModel with the provided choices
func NewEndpointModel(choices []string) endpointModel {
	return endpointModel{choices: choices}
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
			if len(m.choices) > 0 {
				m.choice = m.choices[m.cursor]
				return m, tea.Quit
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m endpointModel) View() string {
	var s strings.Builder

	width, height := getTerminalSize()

	if height < 5 {
		s.WriteString("Terminal is too small to display the options properly.\n")
		return s.String()
	}

	s.WriteString("Please select your module:\n\n")
	s.WriteString("Use hyphen-case for module names with multiple words.\n\n")

	// Calculate start and end indices to display
	start := m.cursor - 1
	if start < 0 {
		start = 0
	}
	end := start + (height - 4) // height - 4 to leave room for header and footer

	// Ensure end index does not exceed the length of choices
	if end > len(m.choices) {
		end = len(m.choices)
	}

	// Ensure start index is not negative
	if start > end {
		start = end
	}

	// Display choices within the range
	for i := start; i < end; i++ {
		if m.cursor == i {
			s.WriteString("▶  ")
		} else {
			s.WriteString("  ")
		}
		choice := m.choices[i]
		if len(choice) > (width - 5) {
			// Truncate long choices if they exceed terminal width
			choice = choice[:width-5] + "..."
		}
		s.WriteString(choice)
		s.WriteString("\n")
	}

	return s.String()
}

// getTerminalSize returns the width and height of the terminal
func getTerminalSize() (width, height int) {
	// Default size in case we can't get the terminal size
	width, height = 80, 16

	// Get terminal size from environment variables if available
	if cols, ok := os.LookupEnv("COLUMNS"); ok {
		width, _ = strconv.Atoi(cols)
	}
	if rows, ok := os.LookupEnv("ROWS"); ok {
		height, _ = strconv.Atoi(rows)
	}

	return width, height
}

// Command to run the endpoint selection
func Command(folders []string) {
	if len(folders) == 0 {
		fmt.Println("No endpoints available.")
		return
	}

	p := tea.NewProgram(NewEndpointModel(folders))

	// Run returns the model as a tea.Model.
	m, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if endpoint, ok := m.(endpointModel); ok && endpoint.choice != "" {
		module := endpoint.choice
		methods.Command(module) // Call method command with selected module
	}
}
