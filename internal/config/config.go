package config

import (
	"log"
	"os"
)

var GlobalConf Config

type Config struct {
	// For all roles, the following fields should be filled
	Role string // svr | cli
	Id   string

	// Info about Role svr
	SvrIp     string
	ProxyPort string

	// Info about Role cli
	ProxyAddr string // <SvrIP>:<ProxyPort>
}

func (c *Config) Load() {
	log.Println("Load configuration")
	c.Role = os.Getenv("HBC_Role")
	c.Id = os.Getenv("HBC_Id")
	c.SvrIp = os.Getenv("HBC_SvrIp")
	c.ProxyPort = os.Getenv("HBC_ProxyPort")
	c.ProxyAddr = c.SvrIp + ":" + c.ProxyPort
}
