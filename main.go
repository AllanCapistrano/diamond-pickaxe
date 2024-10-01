package main

import (
	"fmt"

	"gihub.com/allancapistrano/diamond-pickaxe/cmd"
)

func main() {
	status := cmd.Status()

	fmt.Print(status)
}
