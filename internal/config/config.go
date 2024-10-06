package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kevalsabhani/campus-connect-backend/internal/errors"
)

// Config represents the application configuration
type Config struct {
	Database struct {
		Dsn string `yaml:"dsn" env:"DB_DSN" env-description:"Database DSN"`
	} `yaml:"database"`

	Server struct {
		Port int `yaml:"port" env:"SERVER_PORT" env-description:"Server port"`
	} `yaml:"server"`

	Jwt struct {
		Secret     string `yaml:"secret" env:"JWT_SECRET" env-description:"JWT secret"`
		Expiration int64  `yaml:"expiration" env:"JWT_EXPIRATION" env-description:"JWT expiration"`
	} `yaml:"jwt"`
}

func (c *Config) Validate() error {
	if c.Database.Dsn == "" || c.Server.Port == 0 || c.Jwt.Secret == "" || c.Jwt.Expiration == 0 {
		return errors.ErrEmptyConfig
	}

	if c.Server.Port < 0 || c.Server.Port > 65535 {
		return errors.ErrInvalidPort
	}

	if c.Jwt.Expiration < 0 {
		return errors.ErrInvalidExpiration
	}

	return nil
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
		configFlag := flag.String("config", "config/local.yaml", "path to config file")
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

	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid config: %s", err.Error())
	}

	return &cfg
}
