package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	sslMode := "disable"
	if dbConfig.SSLMode {
		sslMode = "require"
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
		sslMode,
	)

	return connStr
}

func LoadConfig() (*Config, error) {
	isLocal := os.Getenv("IS_LOCAL")

	if isLocal == "true" {
		// Load environment variables from the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		return &Config{
			DBConfig: &DBConfig{
				DBHost:     "localhost",
				DBPort:     "5432",
				DBUser:     "root",
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
