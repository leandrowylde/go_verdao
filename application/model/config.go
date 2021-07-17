package model

import (
	"encoding/json"
	"os"
)

type Config struct {
	URLs  []URI  `json:"urls"`
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

func LoadConfig(file string) (config Config, err error) {
	configFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return
}
