package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/compose-spec/compose-go/types"

	"gopkg.in/yaml.v2"
)

func createAppConfig() types.ServiceConfig {
	return types.ServiceConfig{
		Name: "go-app",
		Environment: types.NewMappingWithEquals([]string{
			"REDIS_HOST=redis-storage",
			"REDIS_PORT=6379",
			"REDIS_PASSWORD=1bfb0c6dbfb9d4f7cd0506ae9e88ffb3",
		}),
		Image: "walenpiotr/url-shortener:1.0.0",
		Ports: []types.ServicePortConfig{
			{
				Target:    8080,
				Published: "8080",
			},
		},
		DependsOn: types.DependsOnConfig{
			"redis-storage": types.ServiceDependency{
				Condition: "service_started",
			},
		},
	}
}

func createRedisConfig() types.ServiceConfig {
	return types.ServiceConfig{
		Name:  "redis-storage",
		Image: "redis:7.0-alpine",
		Ports: []types.ServicePortConfig{
			{
				Target:    6379,
				Published: "6379",
			},
		},
		Restart: "always",
		Command: strings.Split("redis-server --save 20 1 --loglevel warning --requirepass 1bfb0c6dbfb9d4f7cd0506ae9e88ffb3", " "),
	}
}

func main() {
	project := types.Project{
		Services: types.Services{
			createAppConfig(),
			createRedisConfig(),
		},
	}

	b, err := yaml.Marshal(project)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
