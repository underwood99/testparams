package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	digcloudlog "github.com/digautos-library/digCloudLogGo"
)

func running() {
	digcloudlog.DCL_Info(" start running")
	iCount := 0
	text1 := ""
	for {
		iCount++
		if iCount > 10 {
			iCount = 0
			iValue, err := readFile()
			if err != nil {
				digcloudlog.DCL_Error(err.Error())
				iValue = 0
			} else {
				text1 = fmt.Sprintf("running %d", iValue)
				digcloudlog.DCL_Info(text1)
			}
			iValue++
			err = saveFile(iValue)
			if err != nil {
				digcloudlog.DCL_Error("save file error" + err.Error())
				iValue = 0
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func saveFile(id int) error {
	text := fmt.Sprintf("%d", id)
	err := os.WriteFile("/tresss/test1", []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}
func readFile() (int, error) {
	file, err := os.Open("/tresss/test1")
	if err != nil {
		digcloudlog.DCL_Error("not found file /tresss/test1")
	}

	data1, err := io.ReadAll(file)
	if err != nil {
		digcloudlog.DCL_Error("can not read file /tresss/test1")
	}

	i, err := strconv.Atoi(string(data1))

	return i, err
}
