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

	case "tab":
		return KeyPress_Tab(m)

	case "shift+tab":
		return KeyPress_Tab_Shift(m)
	}

	return nil
}
