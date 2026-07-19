package lib

import (
	"os"
	"path/filepath"
)

func GetDir() (string, error) {
	if IsDev {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, ".tmp"), nil
	}

	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}
