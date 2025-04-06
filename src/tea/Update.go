package tea

import (
	"github.com/Kajmany/rapidrule/src/tea/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
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
	case " ": //spacebar
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
	case " ": //spacebar
		m.Mode = normalMode
		return m, cmd
	}
	return m, nil
}
