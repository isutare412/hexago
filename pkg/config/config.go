package config

import "strings"

type Config struct {
	Mode    string         `yaml:"mode"`
	MongoDB *MongoDBConfig `yaml:"mongodb"`
}

func (c *Config) IsProduction() bool {
	mode := strings.ToLower(c.Mode)
	return mode == "production"
}

type MongoDBConfig struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
}
