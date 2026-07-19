package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// get root dir changed
	case lib_GetDir_Res:
		m.states.dir = msg.dir
		if msg.errorMessage != "" {
			setError(m, "root_dir_error", msg.errorMessage)
		} else {
			clearError(m, "root_dir_error")
		}
		return m, bins_GetPath_Cmd(m.states.dir)
	case bins_GetPath_Res:
		m.states.bins.ffmpeg = msg.ffmpeg
		m.states.bins.ytdlp = msg.ytdlp
		if msg.errorMessage != "" {
			setError(m, "bin_error", msg.errorMessage)
		} else {
			clearError(m, "bin_error")
		}
		return m, nil
	case bins_Download_Res:
		m.states.downloading_bin = false
		m.states.progressValue = 0
		if msg.errorMessage != "" {
			setError(m, "bin_error", msg.errorMessage)
			return m, nil
		}
		// no error, clear the error message and re-check if any binaries are missing
		clearError(m, "bin_error")
		return m, bins_GetPath_Cmd(m.states.dir)
	case bins_Download_Progress_Res:
		if !m.states.downloading_bin || msg.percentage >= 1.0 {
			return m, nil
		}
		m.states.progressValue = msg.percentage
		return m, bins_Download_Progress_Cmd(m.progressChannel)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// handle bin downloads
			if m.errorType == "bin_error" {
				// do not start a new download if one is already in progress
				if m.states.downloading_bin {
					return m, nil
				}
				// select which binary to download
				var binary string
				if m.states.bins.ytdlp == "" {
					binary = "yt-dlp"
				} else if m.states.bins.ffmpeg == "" {
					binary = "ffmpeg"
				} else {
					clearError(m, "bin_error")
					return m, nil
				}
				// download the missing binary
				m.states.downloading_bin = true
				m.states.progressValue = 0
				m.errorMessage = fmt.Sprintf("Downloading file \"%s\"", binary)
				return m, tea.Batch(
					bins_Download_Cmd(
						binary,
						m.states.dir,
						m.progressChannel,
					),
					bins_Download_Progress_Cmd(m.progressChannel),
				)
			}
		}
	case tea.WindowSizeMsg:
		m.states.width = msg.Width
		m.states.height = msg.Height
		// return m, tea.ClearScreen
	}

	return m, nil
}

func setError(m *Model, errorType string, errorMessage string) {
	m.errorMessage = errorMessage
	m.errorType = errorType
}
func clearError(m *Model, errorType string) {
	if m.errorType == errorType {
		m.errorMessage = ""
		m.errorType = ""
	}
}
