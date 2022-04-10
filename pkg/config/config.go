package config

type Config struct {
	MongoDB *MongoDBConfig `yaml:"mongodb"`
}

type MongoDBConfig struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
}
