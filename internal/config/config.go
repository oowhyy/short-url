package config

import (
	"os"

	"github.com/oowhyy/short-url/internal/server"
	"github.com/oowhyy/short-url/internal/service"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server      *server.Config  `yaml:"server"`
	StorageType string          `yaml:"storage_type"`
	Service     *service.Config `yaml:"service"`
}

func MustLoadPath(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
