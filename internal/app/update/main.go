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
		cmd := KeyPress(m, msg)
		if cmd != nil {
			return cmd
		}

	case model.Lib_GetDir:
		return handlers.Onchange_lib_GetDir(m, msg)

	case model.Lib_GetSettings:
		return handlers.Onchange_lib_GetSettings(m, msg)

	case model.Bin_GetPath:
		return handlers.Onchange_bin_GetPath(m, msg)

	case model.Bin_Download:
		return handlers.Onchange_bin_download(m, msg)

	case model.UFV_Main_Window_Search_Width:
		m.Main_Window.Search.InputField.Width = msg.Value
		return handlers.Listen_UFV(m.UFV)

	case model.UFV_Bin_Download_ProgressBar_Width:
		m.Bin_Download.ProgressBar.Width = msg.Value
		return handlers.Listen_UFV(m.UFV)
	}

	// send other events to the search input field if it's focused
	if m.Active_layout == model.Layout_Main &&
		m.Main_Window.Active_Panel == model.Active_Panel_Search {
		var cmds []tea.Cmd
		// send events to input field
		if m.Main_Window.Search.InputField.Focused() {
			field, cmd := m.Main_Window.Search.InputField.Update(msg)
			m.Main_Window.Search.InputField = field
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
		// send events to spinner
		if m.Main_Window.Search.Busy {
			spinner, cmd := m.Main_Window.Search.Spinner.Update(msg)
			m.Main_Window.Search.Spinner = spinner
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
		// return batch of commands if any
		if len(cmds) > 0 {
			return tea.Batch(cmds...)
		}
	}

	return nil
}
