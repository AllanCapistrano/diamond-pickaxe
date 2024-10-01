package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

// Executes `git status` command into a specific directory.
func Status() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	vault := os.Getenv("VAULT_PATH")

	commandString := fmt.Sprintf(
		`git -C %s status -s`, vault,
	)

	output, err := exec.Command("/bin/bash", "-c", commandString).Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}
