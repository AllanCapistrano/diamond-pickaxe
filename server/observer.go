package server

import (
	"fmt"
	"time"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

// Check if there are files in the vault to synchronize
func Loop(filePath string, sleep int) { // TODO: Alterar nome do parâmetro 'filePath' para 'vaultPath'
	for {
		if handler.HasFilesToDownload(filePath) {
			fmt.Println("There are files to download!")
		} else { // TODO: Rever lógica, pois talvez isso possa causar conflitos
			if handler.HasFilesToSubmit(filePath) {
				fmt.Println("There are files to submit!")

				submitChanges(filePath)
			} else {
				fmt.Println("There are no files to submit!")
			}
		}

		time.Sleep(time.Duration(sleep) * time.Second)

	}
}

// Submits the local changes
func submitChanges(filePath string) {
	cmd.Add(filePath)

	timeStamp := handler.CurrentTimestampFormatted()

	cmd.Commit(filePath, timeStamp)

	cmd.Push(filePath)
}
