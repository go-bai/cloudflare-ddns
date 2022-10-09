# cloudflare-ddns

NOT SUPPORT NAT
NOT SUPPORT IPv6

// download

```bash
wget https://github.com/go-bai/cloudflare-ddns/blob/main/cloudflare-ddns
```

// init config file

```bash
./cloudflare-ddns init
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
