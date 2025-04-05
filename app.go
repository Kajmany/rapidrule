package main

import (
	"fmt"
	"os"

	"github.com/Kajmany/rapidrule/styles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width      int
	height     int
	statusData table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "i":
			// Scroll up in table
			m.statusData, cmd = m.statusData.Update(msg)
			return m, cmd
		case "j":
			// Scroll down in table
			m.statusData, cmd = m.statusData.Update(msg)
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Update the table width when window size changes
		leftWidth := ((m.width - 2*styles.OuterPadding) * 65) / 100
		m.statusData.SetWidth(leftWidth - 4) // Subtract some space for padding/borders
	}

	// Let the table handle other update events
	m.statusData, cmd = m.statusData.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// Subtract padding for width and height
	innerWidth := m.width - 2*styles.OuterPadding
	innerHeight := m.height - 2*styles.OuterPadding - styles.RibbonHeight

	leftWidth := (innerWidth * 65) / 100
	rightWidth := innerWidth - leftWidth

	// Calculate the available height for the table
	// Account for the title, padding, and borders
	titleHeight := 1   // "Status:" line
	spacingHeight := 2 // Empty lines after title
	borderHeight := 2  // Top and bottom borders
	paddingHeight := 2 // Padding inside the border

	tableHeight := innerHeight - titleHeight - spacingHeight - borderHeight - paddingHeight

	// Adjust table height to fit available space
	m.statusData.SetHeight(tableHeight)

	statusTitle := styles.BoldStyle.Render("Status:")
	tableView := m.statusData.View()

	leftContent := styles.LeftStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(statusTitle + "\n\n" + tableView)

	rightContent := styles.RightStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(styles.BoldStyle.Render("Alerts") + "\n\nDetails, info, or secondary view.\n\nPress 'q' to quit.")

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render("[Q]uit | [I]p | [J]down"))

	return styles.OuterStyle.Render(contentWithRibbon)
}

func main() {

	// Create table columns
	columns := []table.Column{
		{Title: "Local Port", Width: 20},
		{Title: "Peer Addr : Port", Width: 20},
		{Title: "Process", Width: 40},
	}

	// Create dummy rows for testing
	rows := []table.Row{
		{"80", "10.100.168.10:23453", "Nginx"},
		{"3306", "10.100.168.454123", "MySQL"},
		{"22", "10.100.168.454123", "SSHD"},
	}

	// Initialize table with dynamic sizing that will be set properly in View()
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10), // Initial height, will be adjusted in View()
	)

	// Create the proper table.Styles struct instead of using lipgloss.Style
	tableStyles := table.Styles{
		Header: styles.TableStyle.
			Padding(0, 1). // Reduce padding to fit border tighter
			Background(lipgloss.Color("#FFC885")).
			Foreground(lipgloss.Color("#000000")),
		Selected: styles.SelectedStyle.Padding(0, 1), // Reduce padding to fit border tighter
		Cell:     styles.TableStyle.Padding(0, 1),    // Reduce padding to fit border tighter
	}

	t.SetStyles(tableStyles)

	// BubbleTea Logging setup
	// log with log.Println("message")
	log_type := "normal"
	if len(os.Getenv("DEBUG")) > 0 {
		log_type = "debug"
	}
	f, err := tea.LogToFile("debug.log", log_type)
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	// Initialize table with the prepared table
	initialModel := model{
		statusData: t,
	}

	// Use the initialModel that has the table instead of an empty model
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
