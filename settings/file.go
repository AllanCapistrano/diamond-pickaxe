package settings

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

const SETTINGS_FILE_NAME = "diamond-ore.json"
const CONFIG_DIRECTORY_NAME = ".config"
const DIAMOND_PICKAXE_DIRECTORY_NAME = "diamond-pickaxe"

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

// Create the settings file.
func CreateSettingsFile(content Settings) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Couldn't open the user home directory.")
	}

	handler.CreateDirectory(homeDir, CONFIG_DIRECTORY_NAME)

	dirPath := filepath.Join(homeDir, CONFIG_DIRECTORY_NAME)

	handler.CreateDirectory(dirPath, DIAMOND_PICKAXE_DIRECTORY_NAME)

	jsonString, err := json.MarshalIndent(content, "", "    ")
	if err != nil {
		log.Fatal("Couldn't enconde the settings.")
	}

	filePath := filepath.Join(
		homeDir,
		CONFIG_DIRECTORY_NAME,
		DIAMOND_PICKAXE_DIRECTORY_NAME,
		SETTINGS_FILE_NAME,
	)

	os.WriteFile(filePath, jsonString, 0644)
}
