package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type layout struct {
	model   *Model
	strings []string
}

func (l *layout) compile() string {
	return lipgloss.Place(
		l.model.states.width,
		l.model.states.height,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Left, l.strings...),
	)
}
func (l *layout) addTitle(title string) {
	style := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Width(l.model.states.width).MarginBottom(1).Background(lipgloss.Color("#FF0000"))
	l.strings = append(l.strings, style.Render(title))
}
func (l *layout) addProgressBar() {
	l.model.progressBar.ShowPercentage = false
	l.model.progressBar.Width = l.model.states.width
	bar := l.model.progressBar.ViewAs(l.model.states.progressValue)
	l.strings = append(l.strings, bar)
}
func (l *layout) modalBox(text string, modalType string) {
	// set max width to 80
	// subtract 2 from the width to account for the border
	width := min(l.model.states.width-2, 80)

	// prepare the box
	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(width)

	// prepare the content
	content := lipgloss.NewStyle().Padding(1).Render(text)

	// default button is a cancel button
	button := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Width(width).Background(lipgloss.Color("#808080")).Render("Press [ctrl+c] or [q] to exit")

	// determine which button to render based on the modalType
	switch modalType {
	case "download":
		if !l.model.states.downloading_bin {
			button_1 := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#0000FF")).Padding(0, 1).Render("[Enter] to download")
			button_2 := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#FF0000")).Padding(0, 1).Render("[ctrl+c] or [q] to exit")
			gap := lipgloss.NewStyle().Width(
				(width - lipgloss.Width(button_1) - lipgloss.Width(button_2)) / 3,
			).Render()
			text := lipgloss.JoinHorizontal(
				lipgloss.Center,
				button_1,
				gap,
				button_2,
			)
			button = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Width(width).Render(text)
		} else {
			// N.B: here we are using progress bar instead of buttons
			l.model.progressBar.ShowPercentage = false
			l.model.progressBar.Width = width
			button = l.model.progressBar.ViewAs(l.model.states.progressValue)
		}
	}

	// render the box with content and buttons
	rendered := box.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			content,
			button,
		),
	)

	// center the box horizontally and append
	l.strings = append(l.strings, lipgloss.PlaceHorizontal(l.model.states.width, lipgloss.Center, rendered))
}

func handleError(layout *layout) string {
	switch layout.model.errorType {
	case "bin_error":
		layout.modalBox(layout.model.errorMessage, "download")
	default:
		layout.modalBox(layout.model.errorMessage, "cancel")
	}
	return layout.compile()
}

func (m *Model) View() string {
	layout := &layout{model: m}
	layout.addTitle(fmt.Sprintf("GO-TUBE %dX%d @ %s", m.states.width, m.states.height, m.states.dir))

	if m.errorType != "" {
		return handleError(layout)
	}

	layout.addProgressBar()
	return layout.compile()
}
