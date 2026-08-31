# FTN RouterB Architecture

RouterB is the edge/node management service for FTN Cloud. Aether-Core remains the control-plane/orchestration layer.

## Protocol layer

Supported managed adapters: AmneziaWG, WireGuard, Xray, Shadowsocks, Hysteria2, OpenVPN, GRE and IPsec.

Each adapter owns lifecycle/configuration of its upstream protocol implementation; RouterB exposes a common node/API model rather than reimplementing cryptography.

## Clients

- Web: FTN Cloud RouterB console
- Android: FTN Connect/Tunnel client consuming the same API and profile model

## Control plane

`cloud.familytimenet.com` is the configured control-plane domain. Production deployments should terminate TLS at the reverse proxy and restrict the management API to authenticated clients.

## Operational rule

Configuration changes should be explicit, authenticated and auditable. RouterB must not silently alter routing or protocol state.
