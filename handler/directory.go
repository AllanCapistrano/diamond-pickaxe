package handler

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Create a directory in the specified path.
func CreateDirectory(path string, name string) {
	filePath := filepath.Join(path, name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.Mkdir(filePath, os.ModePerm); err != nil {
			log.Fatalf("Couldn't create '%s' directory.", name)
		}
	}
}
