# FTN RouterB Android Client

Android client contract for FTN Connect/Tunnel. The client consumes the RouterB API and presents managed protocol profiles. Native VPN integration should use the platform `VpnService` and the official protocol implementations; this repository does not embed private keys or hard-code server credentials.

Base API: `https://cloud.familytimenet.com/api/v1`

Endpoints: `/info`, `/node`, `/protocols`.
