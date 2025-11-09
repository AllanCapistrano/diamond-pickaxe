package main

import (
	"fmt"
	"log"
	"os"

	"gihub.com/allancapistrano/diamond-pickaxe/server"
	"gihub.com/allancapistrano/diamond-pickaxe/settings"
)

func main() {
	if !settings.CheckSettingsFileExists() {
		settings.CreateSettingsFile(settings.Settings{
			VaultPath:       "",
			VaultRepository: "",
			SyncInSeconds:   30,
		})

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

		fmt.Printf("The settings file was not found. But we created one at '%s'\n", settingsFilePath)
		fmt.Printf("Open the file '%s' to set up the configurations\n", settings.SETTINGS_FILE_NAME)

		os.Exit(1)
	}

	diamondPickaxeSettings := settings.LoadSettingsFile()

	if !settings.IsVaultPathValid(diamondPickaxeSettings.VaultPath) {
		fmt.Printf("The vault path is invalid\n")
		fmt.Printf("Open the file '%s' to set up the configurations\n", settings.SETTINGS_FILE_NAME)
		os.Exit(1)
	}

	if !settings.IsVaultRepositoryValid(diamondPickaxeSettings.VaultRepository) {
		fmt.Printf("The vault repository is invalid\n")
		fmt.Printf("Open the file '%s' to set up the configurations\n", settings.SETTINGS_FILE_NAME)
		os.Exit(1)
	}

	server.Loop(diamondPickaxeSettings.VaultPath, diamondPickaxeSettings.SyncInSeconds)

	// Prevents the program from terminating
	select {}
}
