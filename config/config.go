package config

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ChatID            string
	BotToken          string
	InfluxHost        string
	InfluxPort        string
	InfluxDB          string
	InfluxMeasurement string
	ListenPort        string
}

var configuration Configuration

func init() {
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		log.Panic(err)
	}
}

func Get() Configuration {
	return configuration
}
