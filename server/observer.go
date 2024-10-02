package server

import (
	"fmt"
	"time"

	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

// Check if there are files in the vault to synchronize
func Loop(filePath string, sleep int) {
	for {
		if handler.HasFilesToSync(filePath) {
			fmt.Println("There are files to synchronize!")
		} else {
			fmt.Println("There are no files to synchronize!")
		}

		time.Sleep(time.Duration(sleep) * time.Second)

	}
}
