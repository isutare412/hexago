package config

type Config struct {
	Logger  *LoggerConfig  `yaml:"logger"`
	Timeout *TimeoutConfig `yaml:"timeout"`
	Server  *ServerConfig  `yaml:"server"`
	MongoDB *MongoDBConfig `yaml:"mongodb"`
	Kafka   *KafkaConfig   `yaml:"kafka"`
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
}

type TimeoutConfig struct {
	Startup  int `yaml:"startup"`
	Shutdown int `yaml:"shutdown"`
}

type ServerConfig struct {
	Http *HttpServerConfig `yaml:"http"`
}

type HttpServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
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

type KafkaConfig struct {
	Addrs  []string          `yaml:"addrs"`
	Topics *KafkaTopicConfig `yaml:"topics"`
}

type KafkaTopicConfig struct {
	PaymentRequest *KafkaProducerConfig `yaml:"paymentRequest"`
}

type KafkaProducerConfig struct {
	Topic    string `yaml:"topic"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxRetry int    `yaml:"maxRetry"`
}
