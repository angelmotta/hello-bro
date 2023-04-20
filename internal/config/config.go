package config

import (
	"log"
	"os"
)

var GlobalConf Config

type Config struct {
	// For all roles, the following fields should be filled
	Role string // Role: svr | cli
	Id   string // Id: 0, 1, 2

	// Info about Server
	SvrIp     string // SvrIp: "192.168.1.x"
	ProxyPort string // ProxyPort: "2000"

	// Info about Role cli
	ProxyAddr string // ProxyAddr: <SvrIP>:<ProxyPort>
}

func (c *Config) Load() {
	log.Println("Load configuration")
	c.Role = os.Getenv("HBC_Role")
	c.Id = os.Getenv("HBC_Id")
	c.SvrIp = os.Getenv("HBC_SvrIp")
	c.ProxyPort = os.Getenv("HBC_ProxyPort")
	c.ProxyAddr = c.SvrIp + ":" + c.ProxyPort
}
