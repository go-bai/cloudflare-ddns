package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

type Core struct {
	Config *Config
	API    *cloudflare.API
}

func NewCore(config *Config, api *cloudflare.API) *Core {
	return &Core{
		Config: config,
		API:    api,
	}
}

func (c *Core) Init() {
	defer func() {
		fmt.Println("done")
	}()

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(ex)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(RootPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(RootPath+"cloudflare-ddns", data, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if !c.Config.IsConfigFileExist() {
		err = os.WriteFile(RootPath+c.Config.FileName(), []byte(configDemo), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !c.Config.IsSystemdServiceExist() {
		err := os.WriteFile(RootSystemdPath+"cloudflare-ddns.service", []byte(systemdService), 0644)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func (c *Core) Run() {
	ctx := context.Background()
	zoneID, err := c.API.ZoneIDByName(c.Config.Cloudflare.Domain.Zone)
	if err != nil {
		log.Fatalf("get zone id by name error: %v", err)
	}

	for {
		createOrUpdateRecord(ctx, c, zoneID)
		time.Sleep(30 * time.Second)
	}
}
