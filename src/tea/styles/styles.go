package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Layout constants
const (
	RibbonHeight = 1
	OuterPadding = 1
	LightOrange  = "#FFC885"
	LightGreen   = "#5FAB8F"
	RibbonRed    = "#E67575"
	Lime         = "#CAFE48"
	AppliedGreen = "#7CFC00" // Bright green for applied strategies
	DialogBlue   = "#87CEEB" // Light blue for dialog
)

// Styles used throughout the application
var (
	// OuterStyle is the container style for the entire UI
	OuterStyle = lipgloss.NewStyle()

	NormalModeStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(LightOrange))

	// RibbonStyle is used for the bottom ribbon containing shortcuts
	RibbonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Height(1).
			Background(lipgloss.Color(RibbonRed)).
			Foreground(lipgloss.Color("#000000"))

	// BoldStyle is used to highlight text
	BoldStyle = lipgloss.NewStyle().
			Bold(true)

	// TableStyle is used for table cells and headers
	TableStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(LightOrange)).
			MarginRight(0).
			MarginLeft(0)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Background(lipgloss.Color("#333333")). // Add subtle background to show selection area
			Bold(true)

	// Style for strategies that have been applied
	AppliedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(AppliedGreen)).
			Bold(true)

	// DetailStyle is used for the details section under the table
	DetailStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(LightOrange)).
			MarginRight(0).
			MarginLeft(0)

	PortInfoModeStyle = lipgloss.NewStyle().
				Padding(1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(LightGreen))

	StratModeStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(Lime))

	// Dialog styles for the staging confirmation screen
	DialogStyle = lipgloss.NewStyle().
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(DialogBlue))

	DialogTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(DialogBlue)).
				MarginBottom(1).
				Align(lipgloss.Center)

	DialogOptionStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(DialogBlue))
)
