package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"os"
)

type Mode string

const (
	Development Mode = "dev"
	Production       = "prod"
)

type Config struct {
	Server     ServerConfig   `yaml:"server"`
	DB         DatabaseConfig `yaml:"database"`
	Auth       AuthConfig     `yaml:"auth"`
	ConfigMode Mode           `yaml:"mode"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type AuthConfig struct {
	Length int `yaml:"length"`
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfigFile(path string) (*Config, error) {
	config := NewConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dataBuffer := bytes.NewBuffer(data)
	err = yaml.NewDecoder(dataBuffer).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func SaveConfigFile(path string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	return err
}

func (dc DatabaseConfig) ParseDSN() string {
	return "host=" + dc.Host + " user=" + dc.User + " password=" + dc.Password + " dbname=" + dc.DBName + " sslmode=" + dc.SSLMode
}
