package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	IP   string
	Port int32
}

var AppConfig Config

func SetDafaultAppConfig() {
	AppConfig.IP = "127.0.0.1"
	AppConfig.Port = 8088
}

func ReadConfig() bool {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Successfully Opened config.json")

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &AppConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}(jsonFile)

	return true
}
