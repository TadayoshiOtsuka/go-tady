package config

import (
	"os"
)

type Config struct {
	DB struct {
		Driver   string
		Port     string
		Host     string
		UserName string
		Password string
		DBName   string
	}
}

var AppConfig *Config

func init() {
	AppConfig = newConfig()
}

func newConfig() *Config {
	c := &Config{}
	c.DB.Driver = os.Getenv("DB_DRIVER")
	c.DB.Port = os.Getenv("DB_PORT")
	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.UserName = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")
	c.DB.DBName = os.Getenv("DB_NAME")

	return c
}
