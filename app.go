package main

import (
	"fmt"
	"os"

	localTea "github.com/Kajmany/rapidrule/src/tea"
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/Kajmany/rapidrule/structs"
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

	initialModel.Strats = []structs.Strat{
		{
			Title: "Basic Firewall Setup",
			Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor, nisl eget ultricies ultricies, nunc nisl aliquam nunc, vitae aliquam nisl nunc vitae nisl. Nulla facilisi. Donec euismod, nisl eget ultricies ultricies, nunc nisl aliquam nunc, vitae aliquam nisl nunc vitae nisl. This strategy implements a basic firewall setup with sensible defaults for incoming and outgoing traffic.",
		},
		{
			Title: "Application-Specific Rules",
			Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor. Cras elementum ultrices diam. Maecenas ligula massa, varius a, semper congue, euismod non, mi. This strategy creates specific rules for common applications like web servers, databases, and SSH to ensure only necessary ports are exposed.",
		},
		{
			Title: "Network Segmentation",
			Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras arcu libero, mollis sed, nonummy id, mattis ac, nulla. Curabitur auctor semper nulla. Donec varius orci eget risus. Duis nibh mi, congue eu, accumsan eleifend, sagittis quis, diam. This strategy implements network segmentation to isolate different parts of your network for improved security.",
		},
		{
			Title: "Rate Limiting Protection",
			Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent aliquam, justo convallis luctus rutrum, erat nulla fermentum diam, at nonummy quam ante ac quam. Maecenas urna purus, fermentum id, molestie in, commodo porttitor, felis. This strategy adds rate limiting to protect against DoS attacks and brute force attempts.",
		},
		{
			Title: "Intrusion Detection",
			Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vestibulum accumsan nisl. Nullam eu ante vel est convallis dignissim. Fusce suscipit, wisi nec facilisis facilisis, est dui fermentum leo, quis tempor ligula erat quis odio. This strategy implements basic intrusion detection to alert on suspicious activity.",
		},
	}

	// Use the initialModel that has the table
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
