package main

import (
	"fmt"
	"os"

	"github.com/Kajmany/rapidrule/llm"
	localTea "github.com/Kajmany/rapidrule/src/tea"
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
		{"80", "10.100.168.10:23453", "Nginx"},
		{"3306", "10.100.168.454123", "MySQL"},
		{"22", "10.100.168.454123", "SSHD"},
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

	// Initialize model with the prepared table
	initialModel := localTea.NewModel(t)

	// Use the initialModel that has the table
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	llm.GeminiPortsEval("10.41.80.7:1337 ssh, 10.41.80.7:80 nginx, 10.41.80.7:57 nc")
}
