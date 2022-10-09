package app

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

func createOrUpdateRecord(ctx context.Context, core *Core, zoneID string) {
	ip, err := GetOutboundIP()
	if err != nil {
		fmt.Printf("get outbound ip error: %v\n", err)
		return
	}

	newRecord := cloudflare.DNSRecord{
		TTL:     60,
		Type:    core.Config.Cloudflare.Domain.Type,
		Name:    core.Config.Cloudflare.Domain.Name,
		Content: ip,
	}

	records, err := core.API.DNSRecords(ctx, zoneID, cloudflare.DNSRecord{
		Type: core.Config.Cloudflare.Domain.Type,
		Name: core.Config.Cloudflare.Domain.Name,
	})
	if err != nil {
		log.Fatalf("get dns records error: %v", err)
	}

	if len(records) == 0 {
		msg := fmt.Sprintf("%s >> %s", ip, newRecord.Name)
		if _, err := core.API.CreateDNSRecord(ctx, zoneID, newRecord); err != nil {
			fmt.Printf("unable to create dns record %s, error: %v\n", msg, err)
		} else {
			fmt.Printf("success create dns record: %s\n", msg)
		}
		return
	}

	for _, record := range records {
		if record.Content != ip {
			msg := fmt.Sprintf("%s -> %s", record.Content, newRecord.Content)

			if err = core.API.UpdateDNSRecord(ctx, zoneID, record.ID, newRecord); err != nil {
				fmt.Printf("unable to update record: %s, error: %v\n", msg, err)
			} else {
				fmt.Printf("success update dns record: %s\n", msg)
			}
		}

	}
}
