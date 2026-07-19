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

// wrappers and their types
type lib_GetDir_Res struct {
	dir          string
	errorMessage string
}

func lib_GetDir_Cmd() tea.Msg {
	dir, err := lib.GetDir()
	if err != nil {
		return lib_GetDir_Res{errorMessage: "Failed to get the directory."}
	} else {
		return lib_GetDir_Res{dir: dir}
	}
}

type bins_GetPath_Res struct {
	errorMessage string
	ffmpeg       string
	ytdlp        string
}

func bins_GetPath_Cmd(dir string) tea.Cmd {
	return func() tea.Msg {
		ffmpegPath, ytdlpPath := bins.GetPath(dir)
		res := bins_GetPath_Res{
			ffmpeg: ffmpegPath,
			ytdlp:  ytdlpPath,
		}

		if ytdlpPath == "" {
			res.errorMessage = "\"yt-dlp.exe\" not found in the bin directory."
			return res
		} else if ffmpegPath == "" {
			res.errorMessage = "\"ffmpeg.exe\" not found in the bin directory."
			return res
		} else {
			return res
		}
	}
}

type bins_Download_Res struct {
	errorMessage string
}

func bins_Download_Cmd(binary, dir string, progressCh chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		err := bins.Download(binary, dir, func(downloaded, total int64, percentage float64) {
			progressCh <- bins_Download_Progress_Res{
				percentage: percentage,
			}
		})
		if err == nil {
			return bins_Download_Res{}
		}
		return bins_Download_Res{errorMessage: fmt.Sprintf("Failed to download %s.", binary)}
	}
}

type bins_Download_Progress_Res struct {
	percentage float64
}

func bins_Download_Progress_Cmd(ch chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}
