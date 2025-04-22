package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	return &config, nil
}

type Config struct {
	DB     DatabaseConfig `toml:"database"`
	Server ServerConfig   `toml:"server"`
	JWT    JWTConfig      `toml:"jwt"`
}

type DatabaseConfig struct {
	Host     string `toml:"POSTGRES_HOST"`
	User     string `toml:"POSTGRES_USER"`
	Password string `toml:"POSTGRES_PASSWORD"`
	Database string `toml:"POSTGRES_DB_NAME"`
}

type ServerConfig struct {
	Host string `toml:"HOST"`
	Port int    `toml:"PORT"`
}

type JWTConfig struct {
	ExpiredMinutes int    `toml:"EXPIRED_MINUTES"`
	Secret         string `toml:"JWT_SECRET"`
}

func (c Config) DB_URL() string {
	cfg, _ := LoadConfig("config.toml")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Database)
}

func (c Config) Server_URL() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c Config) JWTSecret() string {
	return c.JWT.Secret
}

func (c Config) JWTMinutes() int {
	return c.JWT.ExpiredMinutes
}
