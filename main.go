package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gihub.com/allancapistrano/diamond-pickaxe/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	vaultPath := os.Getenv("VAULT_PATH")

	sleepInString := os.Getenv("SYNC_IN_SECONDS")

	sleep, err := strconv.Atoi(sleepInString)
	if err != nil {
		logMessage := fmt.Sprintf("Cannot convert '%s' to a number.", sleepInString)

		log.Fatal(logMessage)
	}

	server.Loop(vaultPath, sleep)

	select {}
}
