package main

import (
	"fmt"
	"io"
	"os"

	digcloudlog "github.com/digautos-library/digCloudLogGo"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	fmt.Println("Hello, World!")
	fmt.Println(serverAddress)
	err := digcloudlog.DCL_addLogflare("d0399bff-7fc7-4874-a572-05309021d853", "WoMn49mFQDkh")
	if err != nil {
		fmt.Println(err)
		return
	}
	if serverAddress == "" {
		serverAddress = "no params"
	}
	digcloudlog.DCL_Info(serverAddress)
	// read config file
	file, err := os.Open("/file.txt")
	if err != nil {
		digcloudlog.DCL_Error("didnt find the file")
	}

	data1, err := io.ReadAll(file)
	digcloudlog.DCL_Info(data1)
}
