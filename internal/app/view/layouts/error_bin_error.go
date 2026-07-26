package layouts

import (
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func Error_BinError(m model.Model, width int) []string {
	strArr := []string{}
	strArr = append(strArr, lipgloss.NewStyle().Padding(1).Width(width).Align(lipgloss.Center).Render(m.ErrorMessage))

	// show download progress bar if downloading
	// else show buttons to download or exit
	if m.Bin_Download.Downloading {
		m.Bin_Download.ProgressBar.Width = width
		m.Bin_Download.ProgressBar.ShowPercentage = false
		strArr = append(strArr, m.Bin_Download.ProgressBar.ViewAs(m.Bin_Download.Percentage))
	} else {
		button_1 := lipgloss.NewStyle().Bold(true).Background(m.Theme.Button_BG_main).Padding(0, 1).Render("[Enter] to download")
		button_2 := lipgloss.NewStyle().Bold(true).Background(m.Theme.Button_BG_alt).Padding(0, 1).Render("[ctrl+c] to exit")
		gap := lipgloss.NewStyle().Width(
			(width - lipgloss.Width(button_1) - lipgloss.Width(button_2)) / 3,
		).Render()
		text := lipgloss.JoinHorizontal(
			lipgloss.Center,
			button_1,
			gap,
			button_2,
		)
		strArr = append(strArr, lipgloss.NewStyle().Align(lipgloss.Center).Width(width).Render(text))
	}

	return strArr
}
