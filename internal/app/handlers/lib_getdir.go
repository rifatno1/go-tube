package handlers

import (
	"go-tube/internal/app/model"
	"go-tube/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

func Trigger_lib_GetDir() tea.Cmd {
	return func() tea.Msg {
		dir, err := lib.GetDir()
		if err != nil {
			return model.Lib_GetDir{ErrorMessage: "Failed to get the directory."}
		}
		return model.Lib_GetDir{Dir: dir}
	}
}

func Onchange_lib_GetDir(m *model.Model, msg model.Lib_GetDir) tea.Cmd {
	m.Dir = msg.Dir

	if msg.ErrorMessage != "" {
		setError(m, model.ErrorType_RootDirError, msg.ErrorMessage)
		return nil
	} else {
		clearError(m, model.ErrorType_RootDirError)
		return Trigger_bin_GetPath(m.Dir)
	}
}
