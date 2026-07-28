package update

import (
	"go-tube/internal/app/handlers"
	"go-tube/internal/app/model"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func KeyPress(m *model.Model, msg tea.KeyMsg) tea.Cmd {
	switch strings.ToLower(msg.String()) {
	case "ctrl+c":
		close(m.UFV)
		return tea.Quit

	case "enter":
		if m.ErrorType == model.ErrorType_BinError {
			return handlers.Handle_bin_error(m)
		}
		if m.Active_layout == model.Layout_Main &&
			m.Main_Window.Active_Panel == model.Active_Panel_Search &&
			m.Main_Window.Search.InputField.Focused() {
			m.Main_Window.Search.Busy = true
			m.Main_Window.Search.InputField.Blur()
			return m.Main_Window.Search.Spinner.Tick
		}

	case "tab":
		return KeyPress_Tab(m)

	case "shift+tab":
		return KeyPress_Tab_Shift(m)
	}

	return nil
}
