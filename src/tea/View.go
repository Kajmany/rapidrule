package tea

import (
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/charmbracelet/lipgloss"
)

// View renders the UI
func (m Model) View() string {
	if m.Mode == portInfoMode {
		return m.portInfoView()
	} else {
		return m.normalView()
	}
}

func (m Model) normalView() string {
	// Subtract padding for width and height
	innerWidth := m.Width - 2*styles.OuterPadding
	innerHeight := m.Height - 2*styles.OuterPadding - styles.RibbonHeight

	leftWidth := (innerWidth * 65) / 100
	rightWidth := innerWidth - leftWidth

	// Calculate the available height for the table
	// Account for the title, padding, and borders
	titleHeight := 1   // "Status:" line
	spacingHeight := 2 // Empty lines after title
	borderHeight := 2  // Top and bottom borders
	paddingHeight := 2 // Padding inside the border

	// Reserve space for the detail section (lorem ipsum)
	detailHeight := 6 // Height for the detail section including borders and padding

	// Adjust table height to account for detail section
	tableHeight := innerHeight - titleHeight - spacingHeight - borderHeight - paddingHeight - detailHeight

	// Ensure table height doesn't go below minimum usable size
	if tableHeight < 5 {
		tableHeight = 5
	}

	// Adjust table height to fit available space
	m.StatusData.SetHeight(tableHeight)

	// Adjust table width to fit within the left panel (accounting for borders/padding)
	tableWidth := leftWidth - 4 // 4 = padding + borders
	m.StatusData.SetWidth(tableWidth)

	statusTitle := styles.BoldStyle.Render("Status:")
	tableView := m.StatusData.View()

	// Create the lorem ipsum detail section
	loremText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor."
	detailContent := styles.DetailStyle.
		Width(leftWidth - 4). // Match table width
		Render(styles.BoldStyle.Render("Ai Summary of Network Security Posture:") + "\n" + loremText)

	leftContent := styles.LeftStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(statusTitle + "\n\n" + tableView + "\n\n" + detailContent)

	rightContent := styles.RightStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(styles.BoldStyle.Render("Alerts") + "\n\nDetails, info, or secondary view.\n\nPress 'q' to quit.")

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render("[Q]uit | [↑] Up | [↓] Down | [space] Port Details"))

	return styles.OuterStyle.Render(contentWithRibbon)
}

func (m Model) portInfoView() string {
	// Subtract padding for width and height
	innerWidth := m.Width - 2*styles.OuterPadding
	innerHeight := m.Height - 2*styles.OuterPadding - styles.RibbonHeight

	leftWidth := (innerWidth * 65) / 100
	rightWidth := innerWidth - leftWidth

	// Port info content for the left panel
	aiSummaryTitle := styles.BoldStyle.Render("AI Summary:")
	aiSummaryContent := "\n\nThis port appears to be used by a standard service.\n\nNo unusual activity detected."

	leftContent := styles.LeftStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(aiSummaryTitle + aiSummaryContent)

	// Human summary content for the right panel
	rightContent := styles.RightStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(styles.BoldStyle.Render("Human Summary") + "\n\nWe will write this later\n")

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render("[Q]uit | [space] Normal Mode"))

	return styles.OuterStyle.Render(contentWithRibbon)
}
