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

You can provide `-f` flag that overrides filename

```
go run generate-config/main.go -f docker-compose.test.yml
```

Use `-d` flag to generate dev docker-compose that enables building from local files 
```
go run generate-config/main.go -f docker-compose.dev.yml -d 
```
You can override default image names by using `-ai` - app image and `-ri` - redis image
```
go run generate-config/main.go -ai=walenpiotr/url-shortener:1.1.x -ri=redis:7.x-alpine
```

To specify number of redis shards u `-n` flag (min=1, max=10, default=3)
```
go run generate-config/main.go -n=6
```

**Run app**
```
docker-compose up
```

```
docker-compose -f docker-compose.dev.yml up
```