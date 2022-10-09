# cloudflare-ddns

NOT SUPPORT NAT
NOT SUPPORT IPv6

// download

```bash
wget https://raw.githubusercontent.com/go-bai/cloudflare-ddns/main/cloudflare-ddns
```

// init config file

```bash
chmod +x cloudflare-ddns
./cloudflare-ddns -init
```

// edit config file

```bash
vim /etc/cloudflare-ddns/config.toml
```

// run

```bash
systemctl enable cloudflare-ddns
systemctl start cloudflare-ddns
```
