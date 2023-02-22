# URL Shortener

## Stack
- Go
- Redis
- Docker

## Running app

**Generate docker-compose.yml file using**
```
go run generate-config/main.go
```

You can provide -f flag that overrides filename

```
go run generate-config/main.go -f docker-compose.test.yml
```

**Run app**
```
docker-compose up
```
