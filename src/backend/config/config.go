package config

import (
	"fmt"
	"os"
)

type ConfigItem struct {
	Name, Value, Default string
	Required             bool
}

type ConfigIntItem struct {
	Name     string
	Default  int
	Value    int
	Required bool
}

type Config struct {
	Server struct {
		Env           ConfigItem
		Host          ConfigItem
		Port          ConfigItem
		JWTPrivateKey ConfigItem
	}
	Database struct {
		Host     ConfigItem
		Password ConfigItem
		User     ConfigItem
		Port     ConfigItem
		Name     ConfigItem
	}
}

func PopulateConfigFromEnv(config *Config) error {
	// Server configuration
	config.Server.Env.Default = "development"
	config.Server.Env.Value, _ = os.LookupEnv("ENV")
	config.Server.Env.Required = false
	config.Server.Env.Name = "ENV"
	if config.Server.Env.Value == "" {
		config.Server.Env.Value = config.Server.Env.Default
	}

	config.Server.Host.Default = "localhost"
	config.Server.Host.Value, _ = os.LookupEnv("SERVER_HOST")
	config.Server.Host.Required = false
	config.Server.Host.Name = "SERVER_HOST"
	if config.Server.Host.Value == "" {
		config.Server.Host.Value = config.Server.Host.Default
	}

	config.Server.JWTPrivateKey.Default = "localhost"
	config.Server.JWTPrivateKey.Value, _ = os.LookupEnv("JWT_PRIVATE_KEY")
	config.Server.JWTPrivateKey.Required = false
	config.Server.JWTPrivateKey.Name = "JWT_PRIVATE_KEY"
	if config.Server.JWTPrivateKey.Value == "" {
		config.Server.JWTPrivateKey.Value = config.Server.Host.Default
	}

	config.Server.Port.Default = "8080"
	config.Server.Port.Value, _ = os.LookupEnv("SERVER_PORT")
	config.Server.Port.Required = false
	config.Server.Port.Name = "SERVER_PORT"
	if config.Server.Port.Value == "" {
		config.Server.Port.Value = config.Server.Port.Default
	}

	// Database configuration
	config.Database.Host.Default = "localhost"
	config.Database.Host.Value, _ = os.LookupEnv("DB_HOST")
	config.Database.Host.Required = true
	config.Database.Host.Name = "DB_HOST"
	if config.Database.Host.Value == "" {
		config.Database.Host.Value = config.Database.Host.Default
	}

	config.Database.Password.Default = ""
	config.Database.Password.Value, _ = os.LookupEnv("DB_PASSWORD")
	config.Database.Password.Required = true
	config.Database.Password.Name = "DB_PASSWORD"
	if config.Database.Password.Value == "" {
		config.Database.Password.Value = config.Database.Password.Default
	}

	config.Database.User.Default = ""
	config.Database.User.Value, _ = os.LookupEnv("DB_USER")
	config.Database.User.Required = true
	config.Database.User.Name = "DB_USER"
	if config.Database.User.Value == "" {
		config.Database.User.Value = config.Database.User.Default
	}

	config.Database.Port.Default = "5432"
	config.Database.Port.Value, _ = os.LookupEnv("DB_PORT")
	config.Database.Port.Required = false
	config.Database.Port.Name = "DB_PORT"
	if config.Database.Port.Value == "" {
		config.Database.Port.Value = config.Database.Port.Default
	}

	config.Database.Name.Default = ""
	config.Database.Name.Value, _ = os.LookupEnv("DB_NAME")
	config.Database.Name.Required = false
	config.Database.Name.Name = "DB_NAME"
	if config.Database.Name.Value == "" {
		config.Database.Name.Value = config.Database.Name.Default
	}

	// Validate required values
	for _, item := range []*ConfigItem{
		&config.Database.Host,
		&config.Database.Password,
		&config.Database.User,
	} {
		if item.Required && item.Value == "" {
			return fmt.Errorf("%s is a required environment variable", item.Name)
		}
	}
	return nil
}
