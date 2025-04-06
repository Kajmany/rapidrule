package tea

import (
	"github.com/Kajmany/rapidrule/src/tea/styles"
	"github.com/charmbracelet/lipgloss"
)

// View renders the UI
func (m Model) View() string {
	if m.Mode == portInfoMode {
		return m.portInfoView()
	} else if m.Mode == strategyMode {
		return m.stratView()
	} else if m.Mode == stagingMode {
		return m.stagingView()
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
	overall_text := "No security posture comments at this time."
	if m.AIsummary != "" {
		overall_text = m.AIsummary
	}

	detailContent := styles.DetailStyle.
		Width(leftWidth - 4). // Match table width
		Render(styles.BoldStyle.Render("Ai Summary of Network Security Posture:") + "\n" + overall_text)

	leftContent := styles.NormalModeStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(statusTitle + "\n\n" + tableView + "\n\n" + detailContent)

	alertsTitle := styles.BoldStyle.Render("Alerts:\n")
	alertsContent := ""

	if len(m.Alerts) == 0 {
		alertsContent = "\n\nNo Alerts at this time."
	} else {
		for _, alert := range m.Alerts {
			alertsContent += "\n" + styles.BoldStyle.Render(alert.ShortDesc)
			alertsContent += "\n" + alert.LongDesc + "\n\n"
		}
	}

	// Add help text at the bottom
	if alertsContent == "\n\nNo alerts at this time." {
		alertsContent += "\n\nPress 'q' to quit."
	}

	rightContent := styles.NormalModeStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(alertsTitle + alertsContent)

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render("[Q]uit | [↑] Up | [↓] Down | [<->] Strategy Mode | [space] Port Details"))

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
	if len(m.Ports) > m.StatusData.Cursor() && m.Ports[m.StatusData.Cursor()].Eval != nil {
		aiSummaryContent = "\n\n" + (*(m.Ports[m.StatusData.Cursor()].Eval)).Concerns
	}

	leftContent := styles.PortInfoModeStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(aiSummaryTitle + aiSummaryContent)

	// Human summary content for the right panel
	rightContent := styles.PortInfoModeStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(styles.BoldStyle.Render("Human Summary") + "\n\nWe will write this later\n")

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)
	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render("[Q]uit | [space] Normal Mode"))

	return styles.OuterStyle.Render(contentWithRibbon)
}

func (m Model) stratView() string {
	// Subtract padding for width and height
	innerWidth := m.Width - 2*styles.OuterPadding
	innerHeight := m.Height - 2*styles.OuterPadding - styles.RibbonHeight

	leftWidth := (innerWidth * 35) / 100
	rightWidth := innerWidth - leftWidth

	// Port info content for the left panel
	stratTitle := styles.BoldStyle.Render("Reccomended NFTables Strategies:")
	stratContent := ""

	if len(m.Strats) == 0 {
		stratContent = "\n\nNo Strategies at this time."
	} else {
		// Loop through strategies and highlight the selected one
		for i, strat := range m.Strats {
			// Add newline before each strategy
			if i > 0 {
				stratContent += "\n\n"
			} else {
				stratContent += "\n"
			}

			// Get the title with applied indicator if needed
			title := strat.Title
			if m.AppliedStrats[i] {
				title = "> " + title + " (Applied)"
			}

			// Highlight the selected strategy
			if i == m.StratCursor {
				// Use highlighted style for the selected strategy
				stratContent += styles.SelectedStyle.Render(title)
			} else {
				// Use bold style for non-selected strategies
				if m.AppliedStrats[i] {
					// Use a different style for applied strategies
					stratContent += styles.AppliedStyle.Render(title)
				} else {
					stratContent += styles.BoldStyle.Render(title)
				}
			}

			// Show a preview of the body (first line or so)
			// This keeps the list compact while still providing context
			stratContent += "\n" + truncateString(strat.Body, 50)
		}
	}

	leftContent := styles.StratModeStyle.
		Width(leftWidth).
		Height(innerHeight).
		Render(stratTitle + stratContent)

	// Human summary content for the right panel - show details of selected strategy
	detailTitle := styles.BoldStyle.Render("Strategy Details")
	detailContent := "\n\nSelect a strategy to view details."

	// If we have strategies and a valid cursor, show the details of the selected strategy
	if len(m.Strats) > 0 && m.StratCursor >= 0 && m.StratCursor < len(m.Strats) {
		selectedStrat := m.Strats[m.StratCursor]
		detailContent = "\n\n" + styles.BoldStyle.Render(selectedStrat.Title)

		// Add applied status
		if m.AppliedStrats[m.StratCursor] {
			detailContent += " " + styles.AppliedStyle.Render("(Staged)")
		} else {
			detailContent += " " + styles.BoldStyle.Render("(Not Applied)")
		}

		// Show the strategy details
		detailContent += "\n\n" + selectedStrat.Body
	}

	rightContent := styles.StratModeStyle.
		Width(rightWidth).
		Height(innerHeight).
		Render(detailTitle + detailContent)

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)

	// Update the ribbon to include the spacebar action for applying a strategy
	ribbonMsg := "[Q]uit | [↑] Up | [↓] Down | [<->] Normal Mode | [Enter] Apply Staged Strategies"
	if len(m.Strats) > 0 {
		if m.AppliedStrats[m.StratCursor] {
			ribbonMsg += " | [Space] Remove Staged Strat"
		} else {
			ribbonMsg += " | [Space] Stage Strategy"
		}
	}

	contentWithRibbon := lipgloss.JoinVertical(lipgloss.Top, content, styles.RibbonStyle.Render(ribbonMsg))

	return styles.OuterStyle.Render(contentWithRibbon)
}

func (m Model) stagingView() string {
	// Subtract padding for width and height
	innerWidth := m.Width - 2*styles.OuterPadding
	innerHeight := m.Height - 2*styles.OuterPadding - styles.RibbonHeight

	// Count staged strategies
	stagedCount := 0
	for i := range m.Strats {
		if m.AppliedStrats[i] {
			stagedCount++
		}
	}

	// Create the confirmation dialog
	title := styles.DialogTitleStyle.Render("Confirm Strategy Application")

	// Generate content that shows all staged strategies
	content := "\n\nThe following strategies will be applied:\n\n"

	if stagedCount == 0 {
		content = "\n\nNo strategies have been staged for application.\n\nReturn to strategy view and stage some strategies first."
	} else {
		for i, strat := range m.Strats {
			if m.AppliedStrats[i] {
				content += "• " + styles.BoldStyle.Render(strat.Title) + "\n"
			}
		}
		content += "\n"
	}

	// Add the confirmation options
	if stagedCount > 0 {
		content += styles.DialogOptionStyle.Render("[Y]") + " Yes, apply these strategies\n"
	}
	content += styles.DialogOptionStyle.Render("[N]") + " No, return to strategy view"

	dialogContent := styles.DialogStyle.
		Width(innerWidth - 20). // Make dialog narrower than the full width
		Render(title + content)

	// Center the dialog in the available space
	dialogBox := lipgloss.Place(
		innerWidth,
		innerHeight,
		lipgloss.Center,
		lipgloss.Center,
		dialogContent,
	)

	// Add the ribbon with limited options
	contentWithRibbon := lipgloss.JoinVertical(
		lipgloss.Top,
		dialogBox,
		styles.RibbonStyle.Render("[Y] Apply All | [N] Cancel"),
	)

	return styles.OuterStyle.Render(contentWithRibbon)
}

// Helper function to truncate a string to a specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
