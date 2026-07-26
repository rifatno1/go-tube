package handlers

import (
	"go-tube/internal/app/model"
	"go-tube/internal/bin"

	tea "github.com/charmbracelet/bubbletea"
)

func Trigger_bin_GetPath(dir string) tea.Cmd {
	return func() tea.Msg {
		ffmpegPath, ytdlpPath := bin.GetPath(dir)
		val := model.Bin_GetPath{Ffmpeg: ffmpegPath, Ytdlp: ytdlpPath}
		if ytdlpPath == "" {
			val.ErrorMessage = "\"yt-dlp.exe\" not found in the bin directory."
			return val
		} else if ffmpegPath == "" {
			val.ErrorMessage = "\"ffmpeg.exe\" not found in the bin directory."
			return val
		}
		return val
	}
}

func Onchange_bin_GetPath(m *model.Model, msg model.Bin_GetPath) tea.Cmd {
	m.Bin.Ffmpeg = msg.Ffmpeg
	m.Bin.Ytdlp = msg.Ytdlp
	if msg.ErrorMessage != "" {
		setError(m, model.ErrorType_BinError, msg.ErrorMessage)
	} else {
		clearError(m, model.ErrorType_BinError)
	}
	return nil
}
