package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IP struct {
	Query string
}

func GetOutboundIP() (string, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	var ip IP
	err = json.Unmarshal(body, &ip)
	return ip.Query, err
}
