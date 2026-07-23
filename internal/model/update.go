package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case lib_GetDir:
		return msg.onchange(m)
	case bins_GetPath:
		return msg.onchange(m)
	case bins_Download:
		return msg.onchange(m)
	case bins_Download_Progress:
		return msg.onchange(m)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// handle bin downloads
			if m.errorType == "bin_error" {
				return handle_bin_download(m)
			}
		}
	case tea.WindowSizeMsg:
		m.states.width = msg.Width
		m.states.height = msg.Height
		// return m, tea.ClearScreen
	}

	return m, nil
}
