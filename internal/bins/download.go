package bins

import (
	"errors"
	"go-tube/internal/lib"
	"path/filepath"
)

func Download(file string, baseDir string, onProgress func(downloaded, total int64, percentage float64)) error {
	binDir := filepath.Join(baseDir, "bin")

	fileurl := ""
	filename := ""

	switch file {
	case "ffmpeg":
		filename = "ffmpeg.exe"
		fileurl = "https://github.com/RifatMahmudno-1/ffmpeg-audio-mp3/releases/latest/download/ffmpeg.exe"
	case "yt-dlp":
		fileurl = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"
		filename = "yt-dlp.exe"
	default:
		return errors.New("Invalid file name")
	}

	if err := lib.DownloadFile(fileurl, binDir, filename, onProgress); err != nil {
		return err
	}

	return nil
}
