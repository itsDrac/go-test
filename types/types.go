package types

import "fmt"

type DatabaseConfig struct {
	User             string
	Password         string
	Name             string
	connectionString string
	MaxOpenConns     int
	MaxIdleConns     int
	MaxIdleTime      string
}

func (c *DatabaseConfig) GetConnectionString() string {
	return c.connectionString
}

func (c *DatabaseConfig) SetConnectionString(connectionString string) {
	if connectionString == "" {
		// generate connection string from config
		connectionString = fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", c.User, c.Password, c.Name)
	}
	c.connectionString = connectionString

}
