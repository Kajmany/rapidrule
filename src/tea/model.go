package tea

import (
	"github.com/Kajmany/rapidrule/structs"
	"github.com/charmbracelet/bubbles/table"
)

type inputMode int

const (
	normalMode inputMode = iota
	portInfoMode
)

// Model represents the application state
type Model struct {
	Width      int
	Height     int
	StatusData table.Model
	Mode       inputMode
	Ports      []structs.Port
	Alerts     []structs.Alert
}

// NewModel creates a new model with the given table
func NewModel(t table.Model) Model {
	return Model{
		StatusData: t,
		Mode:       normalMode,
	}
}
