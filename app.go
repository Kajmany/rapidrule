package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
}

var (
	// Add consistent padding to outer container
	outerStyle = lipgloss.NewStyle()

	leftStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFC885"))

	rightStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFC885"))

	// Style for the bottom ribbon for tooltips
	ribbonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Height(1).
			Background(lipgloss.Color("#FFC885")).
			Foreground(lipgloss.Color("#000000"))

	boldStyle = lipgloss.NewStyle().
			Bold(true)
)

// Constants for layout
const (
	ribbonHeight = 1
	outerPadding = 1
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	// Subtract padding for width and height
	innerWidth := m.width - 2*outerPadding
	innerHeight := m.height - 2*outerPadding - ribbonHeight

	leftWidth := (innerWidth * 65) / 100
	rightWidth := innerWidth - leftWidth

	leftContent := leftStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(boldStyle.Render("Status:") + "\n\nHere you can list logs, navigation, etc.")

	rightContent := rightStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(boldStyle.Render("Alerts") + "\n\nDetails, info, or secondary view.\n\nPress 'q' to quit.")

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, ribbonStyle.Render("[Q]uit"))

	return outerStyle.Render(contentWithRibbon)
}

func main() {
	// BubbleTea Logging setup
	// log with log.Println("message")
	log_type := "norma;"
	if len(os.Getenv("DEBUG")) > 0 {
		log_type = "debug"
	}
	f, err := tea.LogToFile("debug.log", log_type)
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
