# 10million

## Setup

1. Clone this repo
1. Install [docker-compose](https://docs.docker.com/compose/install/)
1. Run following commands at the repo root
```sh
# generate ssl cert and keys
openssl req -x509 -newkey rsa:4096 -sha256 -nodes \
    -keyout ./proxy/server.key -out ./proxy/server.crt \
    -subj "/CN=localhost" -days 365
# set up environment variables
cp .env.example .env
```

## Development

```sh
# run all containers
docker-compose up
```

To work on:
- api: work in /api folder, endpoint at https://localhost/api
- web: work in /web folder, endpoint at https://localhost
- db schema: work in /db folder

## Testing

```sh
# api tests
docker-compose exec api go test
```
