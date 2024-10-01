package main

import (
	"fmt"
	"log"
	"os"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	vaultPath := os.Getenv("VAULT_PATH")

	status := cmd.Status(vaultPath)

	fmt.Println(status)

	fmt.Println(handler.HasFilesToSync(vaultPath))
}
