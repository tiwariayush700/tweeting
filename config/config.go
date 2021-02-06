package config

import (
	"encoding/json"
	"flag"
	"github.com/tiwariayush700/tweeting/constants"

	"log"
	"os"
)

const defaultLogLevel = "info"

type Config struct {
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresServer   string `json:"postgres_server"`
	PostgresPort     string `json:"postgres_port"`
	DbName           string `json:"db_name"`
	AuthSecret       string `json:"auth_secret"`
	Port             string `json:"port"`
	LogLevel         string `json:"log_level"`
}

var (
	configuration *Config = nil
	configFile    *string = nil
)

//defined all the required flags
func init() {
	configFile = flag.String(constants.File, constants.DefaultConfig, constants.FileUsage)
}

func ResetConfiguration() {
	configuration = nil
}

func LoadAppConfiguration() {

	if configuration == nil {

		flag.Parse()
		if len(*configFile) == 0 {
			StopService("Mandatory arguments not provided for executing the App")
		}
		configuration = loadConfiguration(*configFile)
	}
}

func loadConfiguration(filename string) *Config {
	if configuration == nil {
		configFile, err := os.Open(filename)
		defer configFile.Close()
		if err != nil {
			StopService(err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		err1 := jsonParser.Decode(&configuration)
		if err1 != nil {
			log.Println("Failed to parse configuration file")
			StopService(err1.Error())
		}
		setDefaultConfig()
	}
	return configuration
}

func GetAppConfiguration() *Config {
	if configuration == nil {
		log.Println("Unable to get the app configuration. Loading freshly. \t")
		LoadAppConfiguration()
	}
	return configuration
}

func StopService(message string) {

	p, _ := os.FindProcess(os.Getpid())
	if err := p.Signal(os.Kill); err != nil {
		log.Fatal("error killing the process while stopping the service")
	}

	log.Fatal(message)
}

func setDefaultConfig() {
	if configuration.LogLevel == "" {
		configuration.LogLevel = defaultLogLevel
	}
}
