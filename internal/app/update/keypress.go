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
		return tea.Quit
	case "enter":
		if m.ErrorType == model.ErrorType_BinError {
			return handlers.Handle_bin_error(m)
		}
	default:
		return nil
	}

	return nil
}
