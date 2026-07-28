package update

import (
	"go-tube/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

func KeyPress_Tab_Shift(m *model.Model) tea.Cmd {
	if m.Active_layout != model.Layout_Main {
		return nil
	}

	switch m.Main_Window.Active_Panel {
	case model.Active_Panel_Search:
		m.Main_Window.Search.InputField.Blur() // remove focus from the search input field
		m.Main_Window.Active_Panel = model.Active_Panel_Playlist
	case model.Active_Panel_Playlist:
		m.Main_Window.Active_Panel = model.Active_Panel_Player
	case model.Active_Panel_Player:
		m.Main_Window.Active_Panel = model.Active_Panel_Musics
	case model.Active_Panel_Musics:
		m.Main_Window.Active_Panel = model.Active_Panel_Search
		if !m.Main_Window.Search.Busy {
			return m.Main_Window.Search.InputField.Focus() // set focus to the search input field
		}
	}

	return nil
}
