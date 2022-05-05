package config

type Config struct {
	Logger  *LoggerConfig  `yaml:"logger"`
	Timeout *TimeoutConfig `yaml:"timeout"`
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
	DonationRequest *KafkaConsumerConfig `yaml:"donationRequest"`
}

type KafkaConsumerConfig struct {
	Topic    string `yaml:"topic"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Group    string `yaml:"group"`
	Timeout  int    `yaml:"timeout"`
}
