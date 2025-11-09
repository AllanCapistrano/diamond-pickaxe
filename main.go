package main

import (
	"gihub.com/allancapistrano/diamond-pickaxe/logger"
	"gihub.com/allancapistrano/diamond-pickaxe/server"
)

const APP_NAME = "diamond-pickaxe"

func main() {
	logger.Init(APP_NAME)

	server.Loop()

	// Prevents the app from terminating
	select {}
}
