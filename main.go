package main

import (
	"fmt"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
	"gihub.com/allancapistrano/diamond-pickaxe/handler"
)

func main() {
	status := cmd.Status()

	fmt.Println(status)

	fmt.Println(handler.HasFilesToSync())
}
