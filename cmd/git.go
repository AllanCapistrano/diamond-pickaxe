package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type git struct {
    Path string
    command func(string, ...any) string
}

// Return the base format for executing git commands.
func create(path string) *git {
    git := &git {
        Path: path,
        command: func(restCommand string, values ...any) string {
            if (len(values) > 0)  {
                restCommand += " %s"
            }

            commandString := fmt.Sprintf(`git -C %s %s`, path, restCommand)

            return fmt.Sprintf(commandString, values...)
        },
    }

    return git
}

// Executes `git status` command into a specific directory.
func Status(path string, parameter string, hasPipeOperator bool, args ...string) string {
	var commandString string
	var output []byte
	var err error

    git := create(path)

    commandString = git.command("status", parameter)

	if hasPipeOperator {
		argsConcatenated := strings.Join(args, " ")

		output, err = exec.Command("/bin/bash", "-c", commandString, "|", argsConcatenated).Output()
	} else {
		output, err = exec.Command("/bin/bash", "-c", commandString).Output()
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

// Executes `git add` command into a specific directory.
func Add(path string) {
    git := create(path)

    commandString := git.command("add", ".")

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not add files.")
	}
}

// Executes `git commit` command into a specific directory.
func Commit(path string, message string) {
	formattedMessage := fmt.Sprintf(`"%s"`, message)

    git := create(path)

    commandString := git.command("commit -m", formattedMessage)

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not create the commit.")
	}
}

// Executes `git push` command into a specific directory.
func Push(path string) {
    git := create(path)

    commandString := git.command("push")

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not push the commit.")
	}
}

// Executes `git fetch` command into a specific directory.
func Fetch(path string) {
    git := create(path)

    commandString := git.command("fetch")

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not fetch the repository.")
	}
}

// Executes `git pull` command into a specific directory.
func Pull(path string) {
    git := create(path)

    commandString := git.command("pull")

	command := exec.Command("/bin/bash", "-c", commandString)

	err := command.Run()
	if err != nil {
		log.Fatal("\nCould not pull the repository.")
	}
}
