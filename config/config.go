package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Port int `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
	MaxConns int32  `yaml:"max_conns"`
}

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = "config/config.yaml"
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}

	if cfg.Server.Port == 0 {
		return nil, errors.New("server.port must be set")
	}
	if cfg.DB.Host == "" || cfg.DB.User == "" || cfg.DB.Name == "" {
		return nil, errors.New("db.{host,user,name} must be set")
	}
	if cfg.DB.MaxConns == 0 {
		cfg.DB.MaxConns = 10
	}
	return &cfg, nil
}
