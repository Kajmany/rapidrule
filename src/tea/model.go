package tea

import (
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
}

// NewModel creates a new model with the given table
func NewModel(t table.Model) Model {
	return Model{
		StatusData: t,
		Mode:       normalMode,
	}
}
