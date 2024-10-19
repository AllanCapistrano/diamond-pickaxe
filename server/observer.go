package server

import (
	"fmt"
	"time"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

// Check if there are files in the vault to synchronize
func Loop(vaultPath string, sleep int) {
	for {
		if handler.HasFilesToDownload(vaultPath) {
			fmt.Println("There are files to download!")

			getChanges(vaultPath)
		} else { // TODO: Rever lógica, pois talvez isso possa causar conflitos
			if handler.HasFilesToSubmit(vaultPath) {
				fmt.Println("There are files to submit!")

				submitChanges(vaultPath)
			} else {
				fmt.Println("There are no files to submit!")
			}
		}

		time.Sleep(time.Duration(sleep) * time.Second)

	}
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
