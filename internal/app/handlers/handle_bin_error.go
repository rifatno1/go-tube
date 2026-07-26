package handlers

import (
	"fmt"
	"go-tube/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

func Handle_bin_error(m *model.Model) tea.Cmd {
	// do not do anything if already downloading
	if m.Bin_Download.Downloading {
		return nil
	}

	// reset the download progress
	m.Bin_Download.Percentage = 0

	// select which binary to download
	var binary string
	if m.Bin.Ytdlp == "" {
		binary = "yt-dlp"
	} else if m.Bin.Ffmpeg == "" {
		binary = "ffmpeg"
	} else {
		// both exist, no need to download
		clearError(m, model.ErrorType_BinError)
		return nil
	}

	// start downloading the binary
	m.Bin_Download.Downloading = true
	m.Bin_Download.ProgressChannel = make(chan model.Bin_Download, 1)
	m.ErrorMessage = fmt.Sprintf("Downloading %s ...", binary)
	return tea.Batch(
		Trigger_bin_download(binary, m.Dir, m.Bin_Download.ProgressChannel),
		Listen_bin_download_progress(m.Bin_Download.ProgressChannel),
	)
}
