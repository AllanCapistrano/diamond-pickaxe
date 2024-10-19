package settings

import (
	"log"
	"os"
	"path/filepath"
)

const SETTINGS_FILE_NAME = "diamond-ore.json"

type Settings struct {
	VaultPath       string `json:"vault_path"`
	VaultRepository string `json:"vault_repository"`
	SyncInSeconds   int    `json:"sync_in_seconds"`
}

// Check if the settings file exists.
func CheckSettingsFileExists() bool {
	foundSettingsFile := true

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Couldn't open the user home directory.")
		foundSettingsFile = false
	}

	filePath := filepath.Join(homeDir, ".config", "diamond-pickaxe", SETTINGS_FILE_NAME)
	file, err := os.Open(filePath)
	if err != nil {
		foundSettingsFile = false
	}

	defer file.Close()

	return foundSettingsFile
}
