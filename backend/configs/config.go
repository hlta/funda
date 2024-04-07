package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	OAuth    OAuthConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// OAuthConfig stores configuration for OAuth, including secrets and keys.
type OAuthConfig struct {
	JWTSecret string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.port", "8080")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&config)
	return
}
