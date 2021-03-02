package config

import (
	"fmt"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JwtSecret string
	Debug     bool
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func (c *DatabaseConfig) ToDsnString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Name,
	)
}
