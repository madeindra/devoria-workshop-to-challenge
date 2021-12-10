package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	App struct {
		Port string
	}
	Database struct {
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
	}
	Bcrypt struct {
		HashCost int
	}
}

func New() *Config {
	c := new(Config)
	c.loadApp()
	c.loadGorm()
	c.loadBcrypt()

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

	connVal := url.Values{}
	connVal.Add("parseTime", "1")
	connVal.Add("loc", "Asia/Jakarta")

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	dsn := fmt.Sprintf("%s?%s", dbConnectionString, connVal.Encode())

	c.Database.DSN = dsn

	c.Database.MaxOpenConnections = int(maxOpenConnections)
	c.Database.MaxIdleConnections = int(maxIdleConnections)

	return c
}

func (c *Config) loadBcrypt() *Config {
	// env value
	hashCost := os.Getenv("BCRYPT_HASH_COST")

	c.Bcrypt.HashCost, _ = strconv.Atoi(hashCost)

	return c
}
