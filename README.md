# Host Scanner API
Host Scanner API is a service that automatically scans your network to discover which hosts are currently online and maps their hostnames to IP addresses. It exposes a simple HTTP API allowing you to retrieve lists of known hosts and their IPs. Scans are performed using Nmap and run on a configurable schedule, with scan results stored to a persistent database. This app is useful for monitoring network inventory or tracking device availability in dynamic environments.

## Running the app:
The app can be run using docker

```
docker pull ghcr.io/durid-ah/host-scanner-api:latest
```

## Configuring the app
The app can be configured using the following environment variables
- `SCANNER_API_HOST` - The IP address the HTTP API server listens on (default: `0.0.0.0`)
- `SCANNER_API_PORT` - The TCP port the API server listens on (default: `8080`)
- `SCANNER_CRON_TAB` - Cron expression defining how often network scans are performed (default: `*/5 * * * *`, which is every 5 minutes)
- `SCANNER_TARGET` - The target network or hosts to scan, in a format supported by Nmap (default: `192.168.1.*`)

## API Endpoints and Pages
- `/docs` - swagger-like page based on the openapi spec for a UI to interract with the API
- `/api/v1/hosts` - List all known hosts and their IP addresses
- `/api/v1/hosts/{hostname}` - Get details (including IP addresses) for a given hostname

