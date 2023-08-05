package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    bool
}

// Expose function to get format string of database connection
func (dbConfig *DBConfig) GetDBConnection() string {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)

	return connStr
}

func LoadConfig() (*Config, error) {
	isLocal := os.Getenv("IS_LOCAL")

	if isLocal == "true" {
		return &Config{
			DBConfig: &DBConfig{
				DBHost:     "localhost",
				DBPort:     "5432",
				DBUser:     "admin",
				DBPassword: "nhin123456",
				DBName:     "lla",
				SSLMode:    false,
			},
		}, nil
	}

	return &Config{
		DBConfig: &DBConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SSLMode:    true,
		},
	}, nil
}
