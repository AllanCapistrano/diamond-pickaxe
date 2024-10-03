package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

// Executes `git status` command into a specific directory.
func Status(path string, useGrep bool) string {
	var commandString string
	var output []byte
	var err error

	if useGrep {
		commandString = fmt.Sprintf(
			`git -C %s status -sb`, path,
		)

		output, err = exec.Command("/bin/bash", "-c", commandString, "|", "grep behind").Output()
	} else {
		commandString = fmt.Sprintf(
			`git -C %s status -s`, path,
		)

		output, err = exec.Command("/bin/bash", "-c", commandString).Output()
	}

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

// Executes `git commit` command into a specific directory.
func Commit(path string, message string) {
	commandString := fmt.Sprintf(
		`git -C %s commit -m "%s"`, path, message,
	)

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not create the commit.")
	}
}

// Executes `git push` command into a specific directory.
func Push(path string) {
	commandString := fmt.Sprintf(
		`git -C %s push`, path,
	)

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not push the commit.")
	}
}

// Executes `git fetch` command into a specific directory.
func Fetch(path string) {
	commandString := fmt.Sprintf(
		`git -C %s fetch`, path,
	)

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not fetch the repository.")
	}
}
