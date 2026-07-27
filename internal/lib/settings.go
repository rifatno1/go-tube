package lib

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Settings struct {
	Quality                 Quality `json:"quality"`
	Theme                   Theme   `json:"theme"`
	Cache_enabled           bool    `json:"cache_enabled"`
	Auto_Download_Next_Song bool    `json:"auto_download_next_song"`
}
type Quality string
type Theme string

const (
	Quality_Low    Quality = "low"
	Quality_Medium Quality = "medium"
	Quality_High   Quality = "high"
	Theme_Light    Theme   = "light"
	Theme_Dark     Theme   = "dark"
)

func getSavedHash(dir string) []byte {
	hashPath := filepath.Join(dir, "data", "settings.hash")
	fileContent, err := os.ReadFile(hashPath)
	if err != nil {
		return nil
	}
	return fileContent
}

func hashMatched(dir string, fileContent []byte) bool {
	oldHash := getSavedHash(dir)
	if oldHash == nil {
		return false
	}
	newHash := sha256.Sum256(fileContent)
	return bytes.Equal(oldHash, newHash[:])
}

func GetSettings(dir string, defaultSettings Settings) (Settings, error) {
	settingsPath := filepath.Join(dir, "data", "settings.json")

	// read settings file
	fileContent, err := os.ReadFile(settingsPath)
	if os.IsNotExist(err) {
		return defaultSettings, nil
	} else if err != nil {
		return defaultSettings, err
	}

	// check if the hash matches the saved hash
	if !hashMatched(dir, fileContent) {
		return defaultSettings, errors.New("Settings file hash does not match the saved hash")
	}

	// unmarshal the settings file into a Settings struct
	userSettings := defaultSettings
	if err := json.Unmarshal(fileContent, &userSettings); err != nil {
		return defaultSettings, err
	}

	// return the user settings
	return userSettings, nil
}

func saveHash(dir string, fileContent []byte) error {
	hash := sha256.Sum256(fileContent)
	hashPath := filepath.Join(dir, "data", "settings.hash")
	return os.WriteFile(hashPath, hash[:], 0644)
}

func SaveSettings(dir string, settings Settings) error {
	settingsPath := filepath.Join(dir, "data", "settings.json")

	// create the data directory if it doesn't exist
	dataDir := filepath.Join(dir, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	// marshal the settings struct into JSON
	fileContent, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	// write the settings file
	if err := os.WriteFile(settingsPath, fileContent, 0644); err != nil {
		return err
	}

	// write hash file
	if err := saveHash(dir, fileContent); err != nil {
		return err
	}

	return nil
}

func GetDefaultSettings() Settings {
	return Settings{
		Quality:                 Quality_Medium,
		Theme:                   Theme_Dark,
		Auto_Download_Next_Song: true,
		Cache_enabled:           true,
	}
}
