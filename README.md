# 10million

## Setup

```sh
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
