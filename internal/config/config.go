package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port      int    `yaml:"port"`
		JWTSecret string `yaml:"jwt_secret"`
	} `yaml:"server"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}

var AppConfig Config

func LoadConfig(path string) error {
	log.Printf("Loading configuration from: %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	return nil
}