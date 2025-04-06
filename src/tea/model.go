package tea

import (
	"log"

	"github.com/Kajmany/rapidrule/structs"
	"github.com/charmbracelet/bubbles/table"
)

type inputMode int

const (
	normalMode inputMode = iota
	portInfoMode
	strategyMode
	stagingMode
)

// Model represents the application state
type Model struct {
	Width       int
	Height      int
	StatusData  table.Model
	Mode        inputMode
	Ports       []structs.Port
	Alerts      []structs.Alert
	AIsummary   string
	Strats      []structs.Strat
	StratCursor int
	// Track whether a strategy has been applied
	AppliedStrats map[int]bool
	// Error message when strategy application fails
	StrategyApplyError string
}

// NewModel creates a new model with the given table
func NewModel(t table.Model) Model {
	return Model{
		StatusData:    t,
		Mode:          normalMode,
		AppliedStrats: make(map[int]bool),
	}
}

// ApplyStrategy is a placeholder function for applying a strategy
// Backend team can hook into this function to implement actual rule application
func (m *Model) ApplyStrategy(stratIndex int) bool {
	if stratIndex < 0 || stratIndex >= len(m.Strats) {
		log.Printf("Cannot apply strategy: index %d out of range", stratIndex)
		return false
	}

	// Get the strategy to apply
	strategy := m.Strats[stratIndex]
	log.Printf("Applying strategy: %s", strategy.Title)

	// Placeholder for actual implementation
	// Backend team should replace this with actual rule application code
	log.Printf("Strategy applied: %s", strategy.Title)

	// Mark this strategy as applied
	m.AppliedStrats[stratIndex] = true

	return true
}

// ApplyAllStagedStrategies applies all the strategies that have been staged
// This is a placeholder for the backend team to implement
func (m *Model) ApplyAllStagedStrategies() bool {
	log.Println("Applying all staged strategies")

	// Count how many strategies are staged
	stagedCount := 0
	for i, strat := range m.Strats {
		if m.AppliedStrats[i] {
			log.Printf("Applying strategy: %s", strat.Title)
			stagedCount++
			// Backend team would implement actual rule application here
		}
	}

	if stagedCount == 0 {
		log.Println("No strategies staged for application")
		return false
	}

	log.Printf("Successfully applied %d strategies", stagedCount)
	return true
}
