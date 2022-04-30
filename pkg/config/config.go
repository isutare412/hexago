package config

type Config struct {
	Logger  *LoggerConfig  `yaml:"logger"`
	MongoDB *MongoDBConfig `yaml:"mongodb"`
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
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
