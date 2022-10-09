package app

import (
	"fmt"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
}
