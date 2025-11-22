package server

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
	"gihub.com/allancapistrano/diamond-pickaxe/settings"
)

// Check if there are files in the vault to synchronize
func Loop() {
	for {
		diamondPickaxeSettings := setup()

		vaultPath := diamondPickaxeSettings.VaultPath
		hasConflictingFiles, files := handler.HasConflictingFiles(vaultPath)

		if hasConflictingFiles {
			fmt.Println("The following files have conflicts:")

			for i := range files {
				fmt.Printf("- %s\n", files[i])
			}
		} else {
			if handler.HasFilesToDownload(vaultPath) {
				fmt.Println("There are files to download!")

				getChanges(vaultPath)
			} else if handler.HasFilesToSubmit(vaultPath) {
				fmt.Println("There are files to submit!")

				submitChanges(vaultPath)
			} else {
				fmt.Println("Nothing to do")
			}
		}

		time.Sleep(time.Duration(diamondPickaxeSettings.SyncInSeconds) * time.Second)
	}
}

// Setup the app settings
func setup() settings.Settings {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Couldn't open the user home directory.")
	}

	settingsFilePath := fmt.Sprintf(
		"%s/%s/%s",
		homeDir,
		settings.CONFIG_DIRECTORY_NAME,
		settings.DIAMOND_PICKAXE_DIRECTORY_NAME,
	)

	if !settings.CheckSettingsFileExists() {
		settings.CreateSettingsFile(settings.Settings{
			VaultPath:     "",
			SyncInSeconds: 30,
		})

		fmt.Printf("The settings file was not found. But we created one at '%s'\n", settingsFilePath)
		fmt.Printf("Open the file '%s' to set up the configurations\n", settings.SETTINGS_FILE_NAME)
		os.Exit(1)
	}

	diamondPickaxeSettings := settings.LoadSettingsFile()

	if !settings.IsVaultPathValid(diamondPickaxeSettings.VaultPath) {
		fmt.Printf("The vault path is invalid\n")
		fmt.Printf("Open the file '%s/%s' to set up the configurations\n", settingsFilePath, settings.SETTINGS_FILE_NAME)
		os.Exit(1)
	}

	return diamondPickaxeSettings
}

// Submits the local changes
func submitChanges(vaultPath string) {
	cmd.Add(vaultPath)

	timeStamp := handler.CurrentTimestampFormatted()

	cmd.Commit(vaultPath, timeStamp)

	cmd.Push(vaultPath)
}

// Get the remote changes
func getChanges(vaultPath string) {
	cmd.Pull(vaultPath)
}
