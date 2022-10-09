package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title      string     `toml:"title" json:"title"`
	Cloudflare Cloudflare `toml:"cloudflare" json:"cloudflare"`
}

type Cloudflare struct {
	CfApiKey   string `toml:"cf_api_key" json:"cf_api_key"`
	CfApiEmail string `toml:"cf_api_email" json:"cf_api_email"`
	Domain     Domain `toml:"domain" json:"domain"`
}

type Domain struct {
	Zone string `json:"zone"`
	Type string `toml:"type" json:"type"`
	Name string `toml:"name" json:"name"`
}

func (c *Config) FileName() string {
	return "config.toml"
}

func (c *Config) IsConfigFileExist() bool {
	_, err := os.Stat(RootPath + c.FileName())
	return err == nil || !os.IsNotExist(err)
}

func (c *Config) IsSystemdServiceExist() bool {
	_, err := os.Stat(RootSystemdPath + "cloudflare-ddns.service")
	return err == nil || !os.IsNotExist(err)
}

func (c *Config) Decode() error {
	if _, err := toml.DecodeFile(RootPath+"config.toml", c); err != nil {
		return err
	}

	strs := strings.Split(c.Cloudflare.Domain.Name, ".")
	if len(strs) >= 2 {
		c.Cloudflare.Domain.Zone = fmt.Sprintf("%s.%s", strs[len(strs)-2], strs[len(strs)-1])
	} else {
		return fmt.Errorf("generate zone failed")
	}
	return nil
}
