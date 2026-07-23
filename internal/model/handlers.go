package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

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

func handle_bin_download(m *Model) (tea.Model, tea.Cmd) {
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
		trigger_bins_Download(
			binary,
			m.states.dir,
			m.progressChannel,
		),
		trigger_bins_Download_Progress(m.progressChannel),
	)
}
