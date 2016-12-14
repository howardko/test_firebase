package util

import (
	"os"
	"encoding/json"
	"fmt"
)

// ConfigStruct for config.json
type ConfigStruct struct {
	Server struct {
		Port     string `json:"port"`
	} `json:"server"`
	Mongo struct {
		Host               string `json:"host"`
		Dbname             string `json:"dbname"`
		TimeoutMillisecond int    `json:"timeout_ms"`
		Auth               struct {
			Enabled                bool   `json:"enabled"`
			Username               string `json:"username"`
			Password               string `json:"password"`
			Mechanism              string `json:"mechanism"`
			AuthenticationDatabase string `json:"authentication_database"`
		} `json:"auth"`
	} `json:"mongo"`
	Query struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"query"`
	Notify struct {
		Server_Key string `json:"server_key"`
	} `json:"notify"`
}

var (
	// Config for env settings
	Config *ConfigStruct
)

// SetupConfig from configPath
func SetupConfig(configPath string) error {
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println("open config file failed. Use defulat config")
	} else {
		Config = &ConfigStruct{}
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(Config); err != nil {
			fmt.Println("parsing config file failed.")
		}
	}
	return err
}
