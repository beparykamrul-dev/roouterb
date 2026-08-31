# FTN RouterB

Multi-protocol router/node control-plane foundation for Family Time Network.

## Included protocols

- AmneziaWG
- WireGuard
- Xray
- Shadowsocks
- Hysteria2
- OpenVPN
- GRE
- IPsec

## Control plane

`cloud.familytimenet.com` → Nginx → RouterB API → protocol/node adapters.

The current release provides a lightweight Go API, health endpoint and web control panel. Protocol adapters are intentionally separated from the control plane so implementations can be added independently.

## Build

```bash
go build -o roouterb ./cmd/router
sudo ./scripts/install.sh
```

## API

- `GET /healthz`
- `GET /api/v1/node`
- `GET /api/v1/protocols`
