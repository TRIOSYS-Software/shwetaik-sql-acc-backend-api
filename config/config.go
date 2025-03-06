package config

import (
	"fmt"
	"os"

	firebird "github.com/flylink888/gorm-firebird"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	ServerIP           string
	ServerPort         string
	DBString           string
	DefinedPreShareKey string
	DefinedKey         string
	AllowedOrigin      string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	config := &Config{}
	config.ServerIP = os.Getenv("SERVER_IP")
	if config.ServerIP == "" {
		config.ServerIP = "localhost"
	}
	config.ServerPort = os.Getenv("SERVER_PORT")
	if config.ServerPort == "" {
		config.ServerPort = "1323"
	}
	config.DBString = os.Getenv("DB_STRING")
	if config.DBString == "" {
		return nil, fmt.Errorf("database connection string is empty")
	}
	config.DefinedPreShareKey = os.Getenv("DEFINED_PRESHAREKEY")
	if config.DefinedPreShareKey == "" {
		return nil, fmt.Errorf("defined presharekey is empty")
	}
	config.DefinedKey = os.Getenv("DEFINED_KEY")
	if config.DefinedKey == "" {
		return nil, fmt.Errorf("defined key is empty")
	}
	config.AllowedOrigin = os.Getenv("ALLOWED_ORIGIN")
	if config.AllowedOrigin == "" {
		config.AllowedOrigin = "http://localhost:3000,http://localhost:1234"
	}
	return config, nil
}

func (c *Config) ConnectDB() (*gorm.DB, error) {
	dsn := c.DBString
	db, err := gorm.Open(firebird.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
