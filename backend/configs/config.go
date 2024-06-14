package configs

import (
	"log"

	"github.com/spf13/viper"
)

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

// CORSConfig holds the configuration for CORS middleware
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string `json:"exposeHeaders"`
	MaxAge           int      `json:"maxAge"`
}

// OAuthConfig stores configuration for OAuth, including secrets and keys.
type OAuthConfig struct {
	JWTSecret string
}

type LogConfig struct {
	Level  string
	Format string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	OAuth    OAuthConfig
	CORS     CORSConfig
	Log      map[string]LogConfig
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.port", "8080")

	viper.AutomaticEnv()

	// Bind environment variables for database config
	viper.BindEnv("database.type", "DATABASE_TYPE")
	viper.BindEnv("database.host", "DATABASE_HOST")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.username", "DATABASE_USERNAME")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.name", "DATABASE_NAME")

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&config)
	return
}
