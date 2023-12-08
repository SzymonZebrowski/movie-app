package configuration

type Config struct {
	Database Database `yaml:"database"`
	Cache Cache `yaml:"cache"`
}

type Database struct {
	Name string `yaml:"name" env:"DB_NAME"`
	User string `yaml:"user" env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Address string `yaml:"address" env:"DB_ADDRESS"`
	Port int `yaml:"port" env:"DB_PORT"`
}

type Cache struct {
	Enabled bool `yaml:"enabled" env:"REDIS_ENABLED"`
	Address string `yaml:"address" env:"REDIS_ADDRESS"`
	Port int `yaml:"port" env:"REDIS_PORT"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
