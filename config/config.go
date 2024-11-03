package config

import (
	"WST_lab1_server/internal/models"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Soap struct {
		PathSoap string `yaml:"pathsoap"`
		PathHttp string `yaml:"pathhttp"`
	} `yaml:"soap"`
	Persons []models.Person `yaml:"persons"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
