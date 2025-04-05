package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Layout constants
const (
	RibbonHeight = 1
	OuterPadding = 1
)

// Styles used throughout the application
var (
	// OuterStyle is the container style for the entire UI
	OuterStyle = lipgloss.NewStyle()

	// LeftStyle is used for the left panel containing the connection list
	LeftStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFC885"))

	// RightStyle is used for the right panel containing alerts
	RightStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFC885"))

	// RibbonStyle is used for the bottom ribbon containing shortcuts
	RibbonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Height(1).
			Background(lipgloss.Color("#FFC885")).
			Foreground(lipgloss.Color("#000000"))

	// BoldStyle is used to highlight text
	BoldStyle = lipgloss.NewStyle().
			Bold(true)

	// TableStyle is used for table cells and headers
	TableStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FFC885")).
			MarginRight(0).
			MarginLeft(0)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Background(lipgloss.Color("#333333")). // Add subtle background to show selection area
			Bold(true)
)
