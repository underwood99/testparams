package main

import (
	"fmt"
	"os"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	fmt.Println("Hello, World!")
	fmt.Println(serverAddress)
}
