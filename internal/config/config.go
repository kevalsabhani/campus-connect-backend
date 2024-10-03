package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents the application configuration
type Config struct {
	Database struct {
		Host     string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
		Port     string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
		Name     string `yaml:"name" env:"DB_NAME" env-description:"Database name"`
		User     string `yaml:"user" env:"DB_USER" env-description:"Database user"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-description:"Database password"`
	} `yaml:"database"`

	Server struct {
		Host string `yaml:"host" env:"SERVER_HOST,HOST" env-description:"Server host"`
		Port string `yaml:"port" env:"SERVER_PORT,PORT" env-description:"Server port"`
	} `yaml:"server"`

	Jwt struct {
		Secret     string `yaml:"secret" env:"JWT_SECRET" env-description:"JWT secret"`
		Expiration int64  `yaml:"expiration" env:"JWT_EXPIRATION" env-description:"JWT expiration"`
	} `yaml:"jwt"`
}

// MustLoad loads the configuration from the file specified by the CONFIG_PATH
// environment variable. If the variable is not set, it checks the -config flag
// and uses its value. If the flag is not set, it logs a fatal error.
// If the file does not exist, it logs a fatal error.
// If the file can not be read, it logs a fatal error.
func MustLoad() *Config {
	var (
		cfg        Config
		configPath string
	)

	// Check the environment variable first
	if configPath = os.Getenv("CONFIG_PATH"); configPath == "" {
		// If the variable is not set, check the -config flag
		configFlag := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *configFlag

		// If the flag is not set, log a fatal error
		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	// Check if the file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}

	// Read the config file
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
