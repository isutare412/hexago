package config

type Config struct {
	Logger  *LoggerConfig  `yaml:"logger"`
	Timeout *TimeoutConfig `yaml:"timeout"`
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
}

type TimeoutConfig struct {
	Startup  int `yaml:"startup"`
	Shutdown int `yaml:"shutdown"`
}

