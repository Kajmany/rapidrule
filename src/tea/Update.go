package tea

import (
	"log"

	"github.com/Kajmany/rapidrule/llm"
	"github.com/Kajmany/rapidrule/scraper"
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/Kajmany/rapidrule/structs"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the model
func (m Model) Init() tea.Cmd {
	var commands []tea.Cmd
	log.Println("init: attempting to scrape ports")
	commands = append(commands, scraper.GetPorts())
	commands = append(commands, scraper.CheckIfRoot())
	commands = append(commands, scraper.CheckTables())
	return tea.Batch(commands...)
}

// Update handles events and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Mode == normalMode {
			return m.updateNormalMode(msg)
		} else if m.Mode == portInfoMode {
			return m.updatePortInfoMode(msg)
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		// Update the table width when window size changes
		leftWidth := ((m.Width - 2*styles.OuterPadding) * 65) / 100
		m.StatusData.SetWidth(leftWidth - 4) // Subtract some space for padding/borders

		// ss has rendered upon us DATA
	case scraper.PortsMsg:
		log.Println("got info for ports")
		for i, port := range msg.Ports {
			log.Printf("%d: %s", i, port.String())
		}
		m.Ports = msg.Ports

		// or not...
	case scraper.PortScrapeError:
		log.Printf("port scrape error message: %s", msg.Err.Error())

<<<<<<< HEAD
	case llm.PortEvalMsg:
		log.Println("got info for port evals")
		for _, eval := range msg.Evals {
			log.Printf("Port %d: %s", eval.Port, eval.String())
			for _, port := range m.Ports {
				if port.Port == eval.Port {
					port.Eval = &eval
					switch eval.Investigate {
					case "No":
						port.LLMRes = structs.Good
					case "Maybe":
						port.LLMRes = structs.Attention
					case "Yes":
						port.LLMRes = structs.Bad
					}
					break
				}
			}
		}

	case llm.PortEvalError:
		log.Printf("port eval error message: %s", msg.Err.Error())
=======
	case scraper.AlertMsg:
		if msg.HasAlert {
			log.Printf("got an alert: %s", msg.Alert.String())
			m.Alerts = append(m.Alerts, msg.Alert)
		} else {
			// Which? who knows?
			log.Printf("alert check came back clear")
		}

	case scraper.AlertError:
		log.Printf("problem trying to check alert status: %s", msg.Err.Error())

>>>>>>> main
	}

	// Let the table handle other update events
	m.StatusData, cmd = m.StatusData.Update(msg)
	return m, cmd
}

func (m Model) updateNormalMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "up", "down": // Handle navigation
		m.StatusData, cmd = m.StatusData.Update(msg)
		return m, cmd
	case " ": // spacebar
		m.Mode = portInfoMode
		return m, cmd
	}

	// Default case - return the model unchanged
	return m, nil
}

func (m Model) updatePortInfoMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case " ": // spacebar
		m.Mode = normalMode
		return m, cmd
	}
	return m, nil
}
