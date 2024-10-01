package handler

import (
	"strings"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
)

// Return true if the vault has files to synchronize. Otherwise, it returns
// false.
func HasFilesToSync(path string) bool {
	filesStatus := cmd.Status(path)

	files := strings.Split(filesStatus, "\n")

	return len(files) > 0
}
