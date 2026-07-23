package model

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		// save the directory to the model
		(&lib_GetDir{}).trigger(),
	)
}

func StartModel() *tea.Program {
	return tea.NewProgram(&Model{
		// Initialize the progress bar
		progressBar:     progress.New(progress.WithDefaultScaledGradient()),
		progressChannel: make(chan tea.Msg, 1),
	}, tea.WithAltScreen())
}
