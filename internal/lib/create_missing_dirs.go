package lib

import (
	"errors"
	"os"
)

func CreateMissingDirs(dir string) error {
	info, err := os.Stat(dir)

	// path exists
	if err == nil {
		if !info.IsDir() {
			return errors.New("Path exists but not a directory")
		} else {
			return nil
		}
	}

	// path does not exist

	// error other than "not exist"
	if !os.IsNotExist(err) {
		return err
	}

	// create the directories
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}
