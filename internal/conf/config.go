package conf

import (
	"errors"
	"os"
)

type Config struct {
	Addr string
	Port string
}

func (c *Config) Init() error {
	if c.Addr = os.Getenv("address"); c.Addr == "" {
		return errors.New("Not found env param 'address'")
	}
	if c.Port = os.Getenv("port"); c.Port == "" {
		return errors.New("Not found env param 'port'")
	}
	return nil
}
