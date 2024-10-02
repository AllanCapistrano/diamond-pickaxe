package server

import (
	"fmt"
	"time"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

// Check if there are files in the vault to synchronize
func Loop(filePath string, sleep int) {
	for {
		if handler.HasFilesToSubmit(filePath) {
			fmt.Println("There are files to synchronize!")

			submitChanges(filePath)
		} else {
			fmt.Println("There are no files to synchronize!")
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
