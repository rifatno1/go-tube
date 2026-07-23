package model

import (
	"fmt"
	"go-tube/internal/bins"
	"go-tube/internal/lib"

	"github.com/charmbracelet/bubbles/progress"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	states          states
	errorMessage    string
	errorType       string
	progressBar     progress.Model
	progressChannel chan tea.Msg
}

type states struct {
	dir           string
	width         int
	height        int
	progressValue float64
	bins          struct {
		ffmpeg string
		ytdlp  string
	}
	downloading_bin bool
}

type lib_GetDir struct {
	dir          string
	errorMessage string
}

func (val *lib_GetDir) trigger() tea.Cmd {
	return func() tea.Msg {
		dir, err := lib.GetDir()
		if err != nil {
			val.errorMessage = "Failed to get the directory."
		} else {
			val.dir = dir
		}
		return *val
	}
}

func (msg *lib_GetDir) onchange(m *Model) (tea.Model, tea.Cmd) {
	m.states.dir = msg.dir
	if msg.errorMessage != "" {
		setError(m, "root_dir_error", msg.errorMessage)
	} else {
		clearError(m, "root_dir_error")
	}
	return m, (&bins_GetPath{}).trigger(m.states.dir)
}

type bins_GetPath struct {
	errorMessage string
	ffmpeg       string
	ytdlp        string
}

func (val *bins_GetPath) trigger(dir string) tea.Cmd {
	return func() tea.Msg {
		ffmpegPath, ytdlpPath := bins.GetPath(dir)
		val.ffmpeg = ffmpegPath
		val.ytdlp = ytdlpPath

		if ytdlpPath == "" {
			val.errorMessage = "\"yt-dlp.exe\" not found in the bin directory."
		} else if ffmpegPath == "" {
			val.errorMessage = "\"ffmpeg.exe\" not found in the bin directory."
		}

		return *val
	}
}

func (msg *bins_GetPath) onchange(m *Model) (tea.Model, tea.Cmd) {
	m.states.bins.ffmpeg = msg.ffmpeg
	m.states.bins.ytdlp = msg.ytdlp
	if msg.errorMessage != "" {
		setError(m, "bin_error", msg.errorMessage)
	} else {
		clearError(m, "bin_error")
	}
	return m, nil
}

type bins_Download struct {
	errorMessage string
}

func (msg *bins_Download) onchange(m *Model) (tea.Model, tea.Cmd) {
	m.states.downloading_bin = false
	m.states.progressValue = 0
	if msg.errorMessage != "" {
		setError(m, "bin_error", msg.errorMessage)
		return m, nil
	}
	// no error, clear the error message and re-check if any binaries are missing
	clearError(m, "bin_error")
	return m, (&bins_GetPath{}).trigger(m.states.dir)
}

func (val *bins_Download) trigger(binary, dir string, progressCh chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		err := bins.Download(binary, dir, func(downloaded, total int64, percentage float64) {
			progressCh <- bins_Download_Progress{
				percentage: percentage,
			}
		})
		if err != nil {
			val.errorMessage = fmt.Sprintf("Failed to download %s.", binary)
		}
		return *val
	}
}

type bins_Download_Progress struct {
	percentage float64
}

func (val *bins_Download_Progress) trigger(ch chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func (msg *bins_Download_Progress) onchange(m *Model) (tea.Model, tea.Cmd) {
	if !m.states.downloading_bin {
		return m, nil
	}
	m.states.progressValue = msg.percentage
	return m, (&bins_Download_Progress{}).trigger(m.progressChannel)
}
