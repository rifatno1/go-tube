package handlers

import (
	"go-tube/internal/app/model"
	"go-tube/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

func Trigger_lib_GetSettings(dir string, defaultSettings lib.Settings) tea.Cmd {
	return func() tea.Msg {
		settings, err := lib.GetSettings(dir, defaultSettings)

		var errorString string
		if err != nil {
			errorString = "Failed to get settings."
		}

		return model.Lib_GetSettings{
			Settings:    settings,
			ErrorString: errorString,
		}
	}
}

func Onchange_lib_GetSettings(m *model.Model, msg model.Lib_GetSettings) tea.Cmd {
	if msg.ErrorString != "" {
		setError(m, model.ErrorType_GetSettingsError, msg.ErrorString)
		return nil
	} else {
		clearError(m, model.ErrorType_GetSettingsError)
		m.Settings = msg.Settings
		return Trigger_bin_GetPath(m.Dir)
	}
}
