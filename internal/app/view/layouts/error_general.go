package layouts

import (
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func Error_General(m model.Model, width int) []string {
	strArr := []string{}
	strArr = append(strArr, lipgloss.NewStyle().Padding(1).Width(width).Align(lipgloss.Center).Render(m.ErrorMessage))
	strArr = append(strArr, lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Width(width).Background(m.Theme.Button_BG_main).Render("Press [ctrl+c] to exit"))

	return strArr
}
