package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	Port   string
}

func LoadConfig() *Config {
	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		log.Fatal("DB_NAME not set in environment variables")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		log.Fatal("DB_USER not set in environment variables")
	}

	dbPass := os.Getenv("POSTGRES_PASSWORD")
	if dbPass == "" {
		log.Fatal("DB_PASSWORD not set in environment variables")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST not set in environment variables")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DBName: dbName,
		DBUser: dbUser,
		DBPass: dbPass,
		DBHost: dbHost,
		DBPort: dbPort,
		Port:   port,
	}
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}
