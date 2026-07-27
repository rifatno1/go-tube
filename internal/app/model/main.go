package model

import (
	"go-tube/internal/lib"

	"github.com/charmbracelet/bubbles/progress"
)

type Model struct {
	Width         int
	Height        int
	Dir           string
	Active_layout layout
	ErrorType     ErrorType
	ErrorMessage  string
	Bin           struct {
		Ffmpeg string
		Ytdlp  string
	}
	Bin_Download struct {
		Downloading     bool
		Percentage      float64
		ProgressBar     progress.Model
		ProgressChannel chan Bin_Download
	}
	Theme       color
	Settings    lib.Settings
	Main_Window main_window
}

type Lib_GetDir struct {
	Dir          string
	ErrorMessage string
}

type Bin_GetPath struct {
	Ffmpeg       string
	Ytdlp        string
	ErrorMessage string
}

type Bin_Download struct {
	Percentage  float64
	ErrorString string
	Completed   bool
}

func GetDefaultBinDownloadProgressBar() progress.Model {
	return progress.New(
		progress.WithDefaultScaledGradient(),
		progress.WithoutPercentage(),
	)
}

type Lib_GetSettings struct {
	Settings    lib.Settings
	ErrorString string
}
