package handler

import (
	"strings"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
)

// Return true if the vault has files to synchronize. Otherwise, it returns
// false.
func HasFilesToSubmit(path string) bool {
	filesStatus := cmd.Status(path, "-s", false)

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

	remoteFilesStatus := cmd.Status(path, "-sb", true, "grep", "behind")

	return strings.Contains(remoteFilesStatus, "[behind 1]")
}

// Return true if there are conflicting files and file names.
// Otherwise, returns false and nil.
func HasConflictingFiles(path string) (bool, []string) {
	output := cmd.LsFiles(path)

	if len(output) != 0 {
		var files []string

		lines := strings.Split(output, "\n")

		filesMap := make(map[string]bool)

		// Avoiding adding files with the same name
		add := func(val string) {
			if !filesMap[val] {
				filesMap[val] = true
			}
		}

		for i := range lines {
			fields := strings.Fields(lines[i])
			filename := fields[len(fields)-1]

			add(filename)
		}

		for i := range filesMap {
			files = append(files, i)
		}

		return true, files
	}

	return false, nil
}
