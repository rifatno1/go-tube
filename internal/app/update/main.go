package update

import (
	"go-tube/internal/app/handlers"
	"go-tube/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

func Update(m *model.Model, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		// return m, tea.ClearScreen

	case tea.KeyMsg:
		return KeyPress(m, msg)

	case model.Lib_GetDir:
		return handlers.Onchange_lib_GetDir(m, msg)

	case model.Lib_GetSettings:
		return handlers.Onchange_lib_GetSettings(m, msg)

	case model.Bin_GetPath:
		return handlers.Onchange_bin_GetPath(m, msg)

	case model.Bin_Download:
		return handlers.Onchange_bin_download(m, msg)
	}

	return nil
}
