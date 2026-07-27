package layouts

import (
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func Main_Search(m model.Model, width int) []string {
	strArr := []string{}

	text_1 := lipgloss.NewStyle().Background(m.Theme.Button_BG_main).Bold(true).Padding(0, 1).Render("Search")
	text_2 := lipgloss.NewStyle().Bold(true).Padding(0, 1).Render("[Enter]")

	w := width - lipgloss.Width(text_2) - lipgloss.Width(text_1)
	if m.Main_Window.Search.InputField.Value() == "" {
		m.Main_Window.Search.InputField.Width = w
	} else {
		m.Main_Window.Search.InputField.Width = w - 4
	}

	bar := m.Main_Window.Search.InputField.View()

	strArr = append(strArr, lipgloss.JoinHorizontal(lipgloss.Left, text_1, bar, text_2))
	return strArr
}
