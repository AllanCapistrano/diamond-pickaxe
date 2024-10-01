package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

// Executes `git status` command into a specific directory.
func Status(path string) string {
	commandString := fmt.Sprintf(
		`git -C %s status -s`, path,
	)

	output, err := exec.Command("/bin/bash", "-c", commandString).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

// Executes `git add` command into a specific directory.
func Add(path string) {
	commandString := fmt.Sprintf(
		`git -C %s add .`, path,
	)

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not add files.")
	}
}
