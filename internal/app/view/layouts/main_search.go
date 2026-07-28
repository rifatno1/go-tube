package layouts

import (
	"go-tube/internal/app/handlers"
	"go-tube/internal/app/model"

	"github.com/charmbracelet/lipgloss"
)

func Main_Search(m model.Model, width int) []string {
	strArr := []string{}

	text_1 := lipgloss.NewStyle().Background(m.Theme.Button_BG_main).Bold(true).Padding(0, 1).Render("Search")
	var text_2 string

	if !m.Main_Window.Search.Busy {
		text_2 = lipgloss.NewStyle().Bold(true).Padding(0, 1).Render("[Enter]")
	} else {
		m.Main_Window.Search.Spinner.Style.Foreground(m.Theme.Text_main)
		spinner := m.Main_Window.Search.Spinner.View()
		text_2 = lipgloss.NewStyle().Bold(true).Padding(0, 1).Render("[Searching" + spinner + "]")
	}

	w := width - lipgloss.Width(text_2) - lipgloss.Width(text_1)
	// 4 is subtracted here because of the width of cursor (1) and width of prompt (3)
	// For some reason, these widths are not counted within the specified width when the field value is not empty
	if m.Main_Window.Search.InputField.Value() != "" {
		w = w - 4
	}

	if m.Main_Window.Search.InputField.Width != w {
		m.Main_Window.Search.InputField.Width = w
		handlers.Push_UFV(m.UFV, model.UFV_Main_Window_Search_Width{Value: w})
	}

	bar := m.Main_Window.Search.InputField.View()

	strArr = append(strArr, lipgloss.JoinHorizontal(lipgloss.Left, text_1, bar, text_2))
	return strArr
}
