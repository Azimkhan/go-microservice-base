package main

import (
	"context"
	"encoding/json"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
)

type Config struct {
	LoggerConfig string `config:"loggerConfig"`
}

func main() {
	conf := loadConfig()
	logger := createLogger(conf.LoggerConfig)
	logger.Info("logger created successfully")
	defer logger.Sync()
}

// Read configuration from file and env variables
func loadConfig() *Config {
	loader := confita.NewLoader(
		file.NewBackend("config.json"),
		env.NewBackend(),
	)
	config := Config{
		"logger.json",
	}
	err := loader.Load(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}

// Read logger configuration from file
func createLogger(configFile string) *zap.Logger {
	zapConf := zap.Config{}
	rawJson, e := ioutil.ReadFile(configFile)
	if e != nil {
		log.Fatal(e)
	}
	e = json.Unmarshal(rawJson, &zapConf)
	if e != nil {
		log.Fatal(e)
	}
	logger, e := zapConf.Build()
	if e != nil {
		log.Fatal(e)
	}
	return logger
}
