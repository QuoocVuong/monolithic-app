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

//package configuration
//
//// Config struct chứa các cấu hình của ứng dụng
//type Config struct {
//	JWTSignerKey string `env:"JWT_SIGNER_KEY"`
//	DBHost       string `env:"DB_HOST"`
//	DBPort       int    `env:"DB_PORT"` // Thay đổi kiểu dữ liệu thành int
//	DBUser       string `env:"DB_USER"`
//	DBPassword   string `env:"DB_PASSWORD"`
//	DBName       string `env:"DB_NAME"`
//	// ... Các cấu hình khác (nếu cần) ...
//}

// LoadConfig tải cấu hình từ biến môi trường
//func LoadConfig() (*Config, error) {
//	cfg := &Config{}
//
//	// Load biến môi trường vào struct
//	if err := loadEnv(cfg); err != nil {
//		return nil, err
//	}
//
//	// Kiểm tra các biến môi trường bắt buộc
//	if cfg.JWTSignerKey == "" {
//		return nil, fmt.Errorf("JWT_SIGNER_KEY environment variable not set")
//	}
//
//	if cfg.DBHost == "" || cfg.DBPort == 0 || cfg.DBUser == "" || cfg.DBName == "" {
//		return nil, fmt.Errorf("database configuration is missing")
//	}
//
//	return cfg, nil
//}
//
//// loadEnv là hàm helper để load biến môi trường vào struct
//func loadEnv(cfg *Config) error {
//	cfg.JWTSignerKey = os.Getenv("JWT_SIGNER_KEY")
//	cfg.DBHost = os.Getenv("DB_HOST")
//	dbPortStr := os.Getenv("DB_PORT")
//	dbPort, err := strconv.Atoi(dbPortStr)
//	if err != nil {
//		return fmt.Errorf("invalid DB_PORT: %w", err)
//	}
//	cfg.DBPort = dbPort
//	cfg.DBUser = os.Getenv("DB_USER")
//	cfg.DBPassword = os.Getenv("DB_PASSWORD")
//	cfg.DBName = os.Getenv("DB_NAME")
//	// ... load các biến môi trường khác
//
//	return nil
//}
//
//// GetDSN trả về DSN để kết nối tới database
//func (c *Config) GetDSN() string {
//	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
//}
