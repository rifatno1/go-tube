package handlers

import (
	"fmt"
	"go-tube/internal/app/model"
	"go-tube/internal/bin"

	tea "github.com/charmbracelet/bubbletea"
)

func Listen_bin_download_progress(channel chan model.Bin_Download) tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-channel
		if !ok {
			return nil
		}
		return msg
	}
}

func Trigger_bin_download(binary, dir string, progressChannel chan model.Bin_Download) tea.Cmd {
	return func() tea.Msg {
		defer close(progressChannel)
		err := bin.Download(binary, dir, func(downloaded, total int64, percentage float64) {
			val := model.Bin_Download{Percentage: percentage}
			// Drain the channel if it already contains an unread older update
			select {
			case <-progressChannel:
			default:
			}
			// Put the newest update into the channel
			select {
			case progressChannel <- val:
			default:
			}
		})
		if err != nil {
			return model.Bin_Download{ErrorString: fmt.Sprintf("Failed to download %s.", binary)}
		}
		return model.Bin_Download{Completed: true}
	}
}

func Onchange_bin_download(m *model.Model, msg model.Bin_Download) tea.Cmd {
	if msg.Completed {
		// download completed successfully, recheck the binary paths
		m.Bin_Download.Downloading = false
		m.Bin_Download.Percentage = 1
		return Trigger_bin_GetPath(m.Dir)
	} else if msg.ErrorString != "" {
		// download failed, set error state
		m.Bin_Download.Downloading = false
		m.Bin_Download.Percentage = 0
		setError(m, model.ErrorType_BinError, msg.ErrorString)
		return nil
	} else {
		m.Bin_Download.Percentage = msg.Percentage
		return Listen_bin_download_progress(m.Bin_Download.ProgressChannel)
	}
}
