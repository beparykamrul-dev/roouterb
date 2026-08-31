#!/usr/bin/env bash
set -euo pipefail
ROOT=/opt/roouterb
install -d "$ROOT/bin" "$ROOT/web"
if [[ -f ./roouterb ]]; then install -m 0755 ./roouterb "$ROOT/bin/roouterb"; else echo 'Build the binary first: go build -o roouterb ./cmd/router'; exit 1; fi
cp -f web/index.html "$ROOT/web/index.html"
install -D -m 0644 systemd/roouterb.service /etc/systemd/system/roouterb.service
systemctl daemon-reload
systemctl enable --now roouterb.service
systemctl --no-pager --full status roouterb.service
