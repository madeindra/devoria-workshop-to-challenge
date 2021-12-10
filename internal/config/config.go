package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	App struct {
		Port string
	}
	Gorm struct {
		DB                 gorm.Dialector
		MaxOpenConnections int
		MaxIdleConnections int
	}
}

func New() *Config {
	c := new(Config)
	c.loadApp()
	c.loadGorm()

	return c
}

func (c *Config) loadApp() *Config {
	// env value
	port := os.Getenv("APP_PORT")

	c.App.Port = port

	return c
}

func (c *Config) loadGorm() *Config {
	// env value
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	maxOpenConnections, _ := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNECTIONS"), 10, 64)
	maxIdleConnections, _ := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONNECTIONS"), 10, 64)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)

	c.Gorm.MaxOpenConnections = int(maxOpenConnections)
	c.Gorm.MaxIdleConnections = int(maxIdleConnections)

	c.Gorm.DB = postgres.Open(dsn)
	return c
}
