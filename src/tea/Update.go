package tea

import (
	"log"

	"github.com/Kajmany/rapidrule/llm"
	"github.com/Kajmany/rapidrule/scraper"
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/Kajmany/rapidrule/structs"
	"github.com/charmbracelet/bubbles/table"
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
	var commands []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Mode == normalMode {
			return m.updateNormalMode(msg)
		} else if m.Mode == strategyMode {
			return m.updateStratMode(msg)
		} else if m.Mode == stagingMode {
			return m.updateStagingMode(msg)
		} else {
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
		var ports_str string
		var ports_strings []string
		for i, port := range msg.Ports {
			log.Printf("%d: %s", i, port.String())
			ports_str += port.ToPrompt()
			ports_strings = append(ports_strings, port.ToPrompt())
		}
		m.Ports = msg.Ports
		var new_rows []table.Row
		for _, port := range m.Ports {
			new_rows = append(new_rows, port.ToRow())
		}
		m.StatusData.SetRows(new_rows)
		commands = append(commands, llm.GetPortEvals(ports_str))
		commands = append(commands, llm.GetTotalEvals(ports_strings))

		// or not...
	case scraper.PortScrapeError:
		log.Printf("port scrape error message: %s", msg.Err.Error())

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

	case llm.PortEvalMsg:
		log.Println("got info for port evals")
		for _, eval := range msg.Evals {
			log.Printf("Port %d: %s", eval.Port, eval.String())
			for i, _ := range m.Ports {
				if m.Ports[i].Port == eval.Port {
					m.Ports[i].Eval = &eval
					switch eval.Investigate {
					case "No":
						m.Ports[i].LLMRes = structs.Good
					case "Maybe":
						m.Ports[i].LLMRes = structs.Attention
					case "Yes":
						m.Ports[i].LLMRes = structs.Bad
					}
					break
				}
			}
		}

	case llm.PortEvalError:
		log.Printf("port eval error message: %s", msg.Err.Error())

	case llm.TotalEvalMsg:
		log.Println("got info for total eval")
		log.Printf("%s", msg.TotalEval.String())

		if msg.TotalEval.Alert == "Red" {
			alert := structs.NewAlert(msg.TotalEval.AlertShort, msg.TotalEval.AlertLong, structs.Red)
			m.Alerts = append(m.Alerts, alert)
		} else if msg.TotalEval.Alert == "Yellow" {
			alert := structs.NewAlert(msg.TotalEval.AlertShort, msg.TotalEval.AlertLong, structs.Yellow)
			m.Alerts = append(m.Alerts, alert)
		}
		m.AIsummary = msg.TotalEval.Overall
		m.View()

	case llm.TotalEvalError:
		log.Printf("total eval error message: %s", msg.Err.Error())
	}

	// Let the table handle other update events
	var cmd tea.Cmd
	m.StatusData, cmd = m.StatusData.Update(msg)
	commands = append(commands, cmd)
	return m, tea.Batch(commands...)
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
	case "left", "right":
		m.Mode = strategyMode
		return m, cmd
	}

	// Default case - return the model unchanged
	return m, nil
}

// for top level strat view
func (m Model) updateStratMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "left", "right":
		m.Mode = normalMode
		return m, cmd
	case "up":
		// Move cursor up, with wraparound
		if len(m.Strats) > 0 {
			m.StratCursor--
			if m.StratCursor < 0 {
				m.StratCursor = len(m.Strats) - 1
			}
		}
		return m, nil
	case "down":
		// Move cursor down, with wraparound
		if len(m.Strats) > 0 {
			m.StratCursor++
			if m.StratCursor >= len(m.Strats) {
				m.StratCursor = 0
			}
		}
		return m, nil
	case " ": // spacebar - toggle staging the currently selected strategy
		if len(m.Strats) > 0 && m.StratCursor >= 0 && m.StratCursor < len(m.Strats) {
			// Toggle the staged state of the current strategy
			m.AppliedStrats[m.StratCursor] = !m.AppliedStrats[m.StratCursor]
		}
		return m, nil
	case "enter": // Enter key - go to staging confirmation screen
		// Switch to staging mode to confirm application of strategies
		m.Mode = stagingMode
		return m, nil
	}
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

// Handle key events in the staging confirmation view
func (m Model) updateStagingMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "n", "N", "esc":
		// Cancel and return to strategy view
		m.Mode = strategyMode
		return m, nil
	case "y", "Y":
		// Apply all staged strategies
		stagedCount := 0
		for i := range m.Strats {
			if m.AppliedStrats[i] {
				stagedCount++
			}
		}

		// Only proceed if there are strategies staged
		if stagedCount > 0 {
			// Apply all staged strategies
			success := m.ApplyAllStagedStrategies()

			if success {
				log.Println("All strategies applied successfully")
				// Exit the application gracefully after applying strategies
				return m, tea.Quit
			} else {
				// Return to strategy view if application failed
				m.Mode = strategyMode
				return m, nil
			}
		} else {
			// If no strategies are staged, just return to strategy view
			m.Mode = strategyMode
			return m, nil
		}
	}

	return m, nil
}
