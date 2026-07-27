package model

import "github.com/charmbracelet/bubbles/textinput"

type main_window struct {
	Active_Panel active_panel
	Search       struct {
		InputField textinput.Model
	}
}

type active_panel struct {
	value int
}

func (a *active_panel) String() int {
	return a.value
}

var (
	Active_Panel_Search   = active_panel{value: 0}
	Active_Panel_Playlist = active_panel{value: 1}
	Active_Panel_Musics   = active_panel{value: 2}
	Active_Panel_Player   = active_panel{value: 3}
)

func GetDefaultMainWindow() main_window {
	mw := main_window{
		Active_Panel: Active_Panel_Search,
	}

	mw.Search.InputField = textinput.New()
	mw.Search.InputField.Blur()
	mw.Search.InputField.CharLimit = 256
	mw.Search.InputField.Placeholder = "Enter YouTube Video/Playlist URL/ID..."
	mw.Search.InputField.Cursor.Blink = true
	mw.Search.InputField.Prompt = " > "
	if mw.Active_Panel == Active_Panel_Search {
		// Currently there is a problem. The cursor doesn't blink now initially.
		// To make the cursor blink, we need to set focus from update and return the event.
		mw.Search.InputField.Focus() // set focus to the search input field
	}

	return mw
}
