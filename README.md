# Distributed Load Balancer

A lightweight HTTP reverse proxy and round-robin load balancer written in Go.

The service accepts a comma-separated list of backend URLs, forwards incoming requests to healthy backends, and retries requests when a backend becomes unavailable.

## Features

- Round-robin request distribution across configured backends
- Reverse proxy forwarding using Go's standard library
- Automatic backend retry on proxy failure
- Backend health tracking with periodic health checks
- Simple Docker and Docker Compose setup for local testing

## How It Works

On startup, the load balancer reads the backend list from the `--backends` flag and creates a reverse proxy for each target.

Request flow:

1. An incoming request is routed to the next healthy backend using round-robin selection.
2. If proxying to that backend fails, the balancer retries the same backend up to 3 times.
3. If it still fails, that backend is marked down for the current routing cycle.
4. The request is retried against another healthy backend up to 3 total attempts.
5. If no healthy backend is available, the service returns `503 Service Unavailable`.

Health checks run every 2 minutes by opening a TCP connection to each backend host.

## Requirements

- Go 1.22+
- Docker and Docker Compose, if running the containerized example

## Run Locally

Start the load balancer with one or more backend servers:

```bash
go run . --backends "http://localhost:8081,http://localhost:8082" --port 3030
```

Build a binary:

```bash
go build -o lb .
./lb --backends "http://localhost:8081,http://localhost:8082" --port 3030
```

## CLI Options

- `--backends`: Comma-separated backend URLs. Required.
- `--port`: Port for the load balancer to listen on. Default: `3030`

Example:

```bash
./lb --backends "http://web1:80,http://web2:80,http://web3:80" --port 3030
```

## Run With Docker Compose

The repository includes a Compose setup with three sample backend containers and one load balancer container.

Start the stack:

```bash
docker compose up --build
```

Stop the stack:

```bash
docker compose down
```

Once started, send requests to:

```text
http://localhost:3030
```

Repeated requests should rotate across `web1`, `web2`, and `web3` while they remain healthy.

## Project Structure

```text
.
├── main.go             # Load balancer implementation
├── go.mod              # Go module definition
├── dockerfile          # Multi-stage Docker build
├── docker-compose.yml  # Local multi-container demo
└── README.md           # Project documentation
```

## Notes

- This project handles HTTP load balancing only.
- Health checks use TCP connectivity, not HTTP status validation.
- Backend state is stored in memory and resets when the process restarts.
- There are currently no automated tests in the repository.
