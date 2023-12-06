package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Config struct {
	GrpcIP          string
	GrpcPort        int32
	MongodbIP       string
	MongodbPort     int32
	MongodbUser     string
	MongodbPassword string
}

var AppConfig Config

func SetDafaultAppConfig() {
	AppConfig.GrpcIP = "127.0.0.1"
	AppConfig.GrpcPort = 8088
	AppConfig.MongodbIP = "127.0.0.1"
	AppConfig.MongodbPort = 27017
}

func ReadConfig() bool {
	jsonFile, err := os.Open("config_dev.json")

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

	fmt.Println("GrpcIP:" + AppConfig.GrpcIP)
	fmt.Println("GrpcPort:" + strconv.Itoa(int(AppConfig.GrpcPort)))
	fmt.Println("MongodbIP:" + AppConfig.MongodbIP)
	fmt.Println("MongodbPort:" + strconv.Itoa(int(AppConfig.MongodbPort)))

	return true
}
