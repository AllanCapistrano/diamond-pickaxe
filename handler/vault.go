package handler

import (
	"strings"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
)

// Return true if the vault has files to synchronize. Otherwise, it returns
// false.
func HasFilesToSubmit(path string) bool {
	filesStatus := cmd.Status(path, false)

	files := strings.Split(filesStatus, "\n")

	filteredFiles := []string{}
	for _, file := range files {
		if file != "" {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return len(filteredFiles) > 0
}

// Return true if there are files in the remote vault that need to be downloaded.
// Otherwise, it returns false.
func HasFilesToDownload(path string) bool {
	cmd.Fetch(path)

	remoteFilesStatus := cmd.Status(path, true) // TODO: Está dando erro ao usar o '| grep behind'

	files := strings.Split(remoteFilesStatus, "\n")

	filteredFiles := []string{}
	for _, file := range files {
		if file != "" {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return len(filteredFiles) > 0
}