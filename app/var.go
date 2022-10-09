package app

var RootPath = "/etc/cloudflare-ddns/"
var RootSystemdPath = "/etc/systemd/system/"

var configDemo = `title = "Oracle JP"

[cloudflare]
cf_api_key = "xxxx"
cf_api_email = "xxx@xx.com"

[cloudflare.domain]
type = "A"
name = "test.xxx.xxx"`

var systemdService = `
[Unit]
Description=cloudflare-ddns service
After=network.target nss-lookup.target
Wants=network.target

[Service]
User=root
Group=root
Type=simple
LimitAS=infinity
LimitRSS=infinity
LimitCORE=infinity
LimitNOFILE=999999
WorkingDirectory=/etc/cloudflare-ddns/
ExecStart=/etc/cloudflare-ddns/cloudflare-ddns -run
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target`
