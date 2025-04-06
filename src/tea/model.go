package tea

import (
	"github.com/charmbracelet/bubbles/table"
)

// Model represents the application state
type Model struct {
	Width      int
	Height     int
	StatusData table.Model
}

// NewModel creates a new model with the given table
func NewModel(t table.Model) Model {
	return Model{
		StatusData: t,
	}
}
