# 10million

## Setup

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

## Testing

```sh
# api tests
docker-compose exec api go test
```
