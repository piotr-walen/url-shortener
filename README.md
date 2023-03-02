# URL Shortener

This is a URL shortener application that allows users to create custom short URLs with namespaces and aliases. 
The custom short URLs are created in the following format: `http://url-shortener-domain/${namespace}/${alias}`.

The app backend is built using Go and it uses Redis as datastore. 
It also has a custom sharding mechanism using the hashring library from http://github.com/serialx/hashring. 
Additionally, there is a custom script that generates a Docker Compose configuration file that allows for the spawning of multiple Redis shards.

The browser client is built using Vite and React.


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