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

Use `-db` flag to generate dev docker-compose that enables building docker images from local files 
```
go run generate-config/main.go -f docker-compose.build.yml -db
```

Use `-dh` flag to generate dev docker-compose that enables hot module replacement
```
go run generate-config/main.go -f docker-compose.hmr.yml -dh
```

You can override default image names by using `-ai` - app image and `-ri` - redis image
```
go run generate-config/main.go -ai=walenpiotr/url-shortener:1.1.x -ri=redis:7.x-alpine
```

To specify number of redis shards u `-n` flag (min=1)
```
go run generate-config/main.go -n=6
```

**Run app**
Using hosted images
```
docker-compose up
```

With building building docker images from local files
```
docker-compose -f docker-compose.hmr.yml up --build
```

With HMR
```
docker-compose -f docker-compose.hmr.yml up --build
```