package configuration

import (
	"fmt"
	"os"
)

type Config struct {
	JWTSignerKey string
	DBConfig     struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	// ... Các cấu hình khác (nếu cần) ...
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		JWTSignerKey: os.Getenv("JWT_SIGNER_KEY"),
		DBConfig: struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		// ... Các cấu hình khác ...
	}

	if cfg.JWTSignerKey == "" {
		return nil, fmt.Errorf("JWT_SIGNER_KEY environment variable not set")
	}

	// Kiểm tra các biến môi trường database
	if cfg.DBConfig.Host == "" || cfg.DBConfig.Port == "" || cfg.DBConfig.User == "" || cfg.DBConfig.DBName == "" {
		return nil, fmt.Errorf("database configuration is missing")
	}

	return cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBConfig.User, c.DBConfig.Password, c.DBConfig.Host, c.DBConfig.Port, c.DBConfig.DBName)
}
