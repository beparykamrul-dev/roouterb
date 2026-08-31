# FTN RouterB

Multi-protocol router/node control-plane service for Family Time Network.

## Managed protocol families

AmneziaWG, WireGuard, Xray, Shadowsocks, Hysteria2, OpenVPN, GRE and IPsec.

RouterB is the node/edge service. Aether-Core is the orchestration/control-plane layer. The configured cloud endpoint is `cloud.familytimenet.com`.

## Repository layout

- `cmd/routerb` — Go RouterB service
- `cmd/router` — existing router entrypoint
- `web` — web console
- `android` — Android client/API contract
- `config` — protocol configuration
- `deploy` — Docker/Nginx deployment
- `systemd` — hardened service unit
- `scripts` — installation helpers
- `docs` — architecture documentation

## Local build

```bash
go test ./...
go build -trimpath -o roouterb ./cmd/routerb
ROUTERB_ADDR=127.0.0.1:8080 ./roouterb
```

## API

- `GET /health`
- `GET /healthz`
- `GET /api/v1/info`
- `GET /api/v1/node`
- `GET /api/v1/protocols`

## Deployment

Docker and systemd definitions are included. Put TLS/authentication at the production reverse proxy and do not commit credentials, private keys, or protocol secrets.
