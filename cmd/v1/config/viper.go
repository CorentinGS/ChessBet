package config

import (
	"log/slog"

	"github.com/corentings/chessbet/infrastructures"
	"github.com/corentings/chessbet/pkg/oauth"
	"github.com/spf13/viper"
)

// Config holds the configuration for the application.

type Config struct {
	Auth   AuthConfigStruct
	Server ServerConfig
	Log    LogConfig
	Pgx    infrastructures.PgxConfig
}

// AuthConfigStruct is the auth configuration.
type AuthConfigStruct struct {
	Discord       oauth.DiscordConfig
	Github        oauth.GithubConfig
	JWTSecret     string
	JWTHeaderLen  int
	JWTExpiration int
}

// ServerConfig holds the configuration for the server.
type ServerConfig struct {
	Port        string
	AppVersion  string
	JaegerURL   string
	Host        string
	FrontendURL string
	LogLevel    string
}

// LogConfig holds the configuration for the logger.
type LogConfig struct {
	Level string
}

func (logConfig *LogConfig) GetSlogLevel() slog.Level {
	switch logConfig.Level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// LoadConfig loads the configuration from a file.
func LoadConfig(filename string) (*Config, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
