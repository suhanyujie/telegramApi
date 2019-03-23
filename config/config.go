package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var ApiList = []string{
	"getUpdates",
}

type Config struct {
	BaseUrl     string      `json:"baseurl"`
	ApiToken    string      `json:"apitoken"`
	TelegramBot TelegramBot `json:"telegramBot"`
}

type TelegramBot struct {
	ChannelRoom string `json:"channelRoom"`
}

var ConfigData Config
var configPath *string

func init() {
	configPath = flag.String("c", "./config.json", "set config path:-c=config.json")
	flag.Parse()
	log.Println(configPath)
	ParseConfig(*configPath)
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	content, err := ioutil.ReadAll(file)
	err = json.Unmarshal(content, &ConfigData)
	if err != nil {
		log.Fatalln(err)
	}

	return &ConfigData, nil
}
