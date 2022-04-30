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
	Addrs             []string `yaml:"addrs"`
	Database          string   `yaml:"database"`
	AuthSource        string   `yaml:"authSource"`
	Username          string   `yaml:"username"`
	Password          string   `yaml:"password"`
	HeartbeatInterval int      `yaml:"heartbeatInterval"`
	MaxConnectionPool int      `yaml:"maxConnectionPool"`
}
