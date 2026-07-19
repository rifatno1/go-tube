package lib

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type DownloadProgressWriter struct {
	Total      int64
	Downloaded int64
	Percentage float64
	OnProgress func(downloaded, total int64, percentage float64)
}

func (pw *DownloadProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Downloaded += int64(n)
	if pw.Total > 0 {
		pw.Percentage = float64(pw.Downloaded) / float64(pw.Total)
	}
	if pw.OnProgress != nil {
		pw.OnProgress(pw.Downloaded, pw.Total, pw.Percentage)
	}
	return n, nil
}

func DownloadFile(url string, dir string, filename string, onProgress func(downloaded, total int64, percentage float64)) error {
	if err := CreateMissingDirs(dir); err != nil {
		return err
	}

	filePath := filepath.Join(dir, filename)
	info, err := os.Stat(filePath)
	// directory or file already exists
	if err == nil {
		if info.IsDir() {
			return errors.New("A directory with the same name already exists")
		} else {
			return errors.New("File already exists")
		}
	}
	// other error than "not exist"
	if !os.IsNotExist(err) {
		return err
	}

	// download the file
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to download file: " + resp.Status)
	}

	// create a progress writer to track the download progress
	pw := &DownloadProgressWriter{
		Total:      resp.ContentLength,
		OnProgress: onProgress,
	}

	// create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	var downloadSuccess bool
	// close it not closed
	defer func() {
		out.Close()
		if !downloadSuccess {
			os.Remove(filePath)
		}
	}()

	// write to file and track progress
	_, err = io.Copy(io.MultiWriter(out, pw), resp.Body)
	if err != nil {
		return err
	}
	downloadSuccess = true

	return nil
}
