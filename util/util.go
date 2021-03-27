package util

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration server Config
type Configuration struct {
	Version      string
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// Config object for app
var Config Configuration

// Logger for app
var Logger *log.Logger

func init() {
	loadConfig()
	initLogger()
}

func initLogger() {
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open Config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// Info message
func Info(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

// Danger message
func Danger(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

// Warning message
func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}
