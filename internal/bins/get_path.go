package bins

import (
	"os"
	"path/filepath"
)

func GetPath(dir string) (string, string) {
	binDir := filepath.Join(dir, "bin")

	ffmpegPath := ""
	ytdlpPath := ""

	if info, err := os.Stat(filepath.Join(binDir, "ffmpeg.exe")); err == nil && !info.IsDir() {
		ffmpegPath = filepath.Join(binDir, "ffmpeg.exe")
	}

	if info, err := os.Stat(filepath.Join(binDir, "yt-dlp.exe")); err == nil && !info.IsDir() {
		ytdlpPath = filepath.Join(binDir, "yt-dlp.exe")
	}

	return ffmpegPath, ytdlpPath
}
