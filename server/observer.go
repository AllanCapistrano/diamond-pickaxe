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
		if handler.HasFilesToSync(filePath) {
			fmt.Println("There are files to synchronize!")

			synchronizeFiles(filePath)
		} else {
			fmt.Println("There are no files to synchronize!")
		}

		time.Sleep(time.Duration(sleep) * time.Second)

	}
}

// Synchronize the files
func synchronizeFiles(filePath string) {
	cmd.Add(filePath)

	timeStamp := handler.CurrentTimestampFormatted()

	cmd.Commit(filePath, timeStamp)
}
