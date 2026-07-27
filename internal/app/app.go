package app

import (
	appInit "go-tube/internal/app/init"
	appModel "go-tube/internal/app/model"
	appUpdate "go-tube/internal/app/update"
	appView "go-tube/internal/app/view"
	"go-tube/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

type app struct {
	model appModel.Model
}

func (m app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := appUpdate.Update(&m.model, msg)
	return m, cmd
}

func (m app) View() string {
	return appView.View(m.model)
}

func (m app) Init() tea.Cmd {
	return appInit.Init(m.model)
}

func StartApp() *tea.Program {
	defaultApp := app{}

	// Downaload progress bar initialization
	defaultApp.model.Bin_Download.ProgressBar = appModel.GetDefaultBinDownloadProgressBar()

	// Set the default layout to main layout
	defaultApp.model.Active_layout = appModel.Layout_Main

	// set default settings
	defaultApp.model.Settings = lib.GetDefaultSettings()

	// set default theme
	switch defaultApp.model.Settings.Theme {
	case lib.Theme_Light:
		defaultApp.model.Theme = appModel.Theme_Light
	default:
		defaultApp.model.Theme = appModel.Theme_Dark
	}

	// set default main window
	defaultApp.model.Main_Window = appModel.GetDefaultMainWindow()

	return tea.NewProgram(&defaultApp, tea.WithAltScreen())
}
