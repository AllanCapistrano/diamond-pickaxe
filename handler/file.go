package handler

import (
	"strings"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
)

// Return true if the vault has files to synchronize. Otherwise, it returns
// false.
func HasFilesToSync() bool {
	files_status := cmd.Status()

	files := strings.Split(files_status, "\n")

	return len(files) > 0
}
