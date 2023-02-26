package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/compose-spec/compose-go/types"

	"gopkg.in/yaml.v2"
)

var FILENAME *string
var DEV_HMR *bool
var DEV_BUILD *bool
var REDIS_IMAGE *string
var APP_IMAGE *string
var NUM_SHARDS *int
var CLIENT_IMAGE *string

func parseFlags() {
	FILENAME = flag.String("f", "docker-compose.yml", "a filename of generated docker compose file")
	DEV_BUILD = flag.Bool("db", false, "generate dev file that allows building images from local files")
	DEV_HMR = flag.Bool("dh", false, "generate dev file that allows hot reloading")
	REDIS_IMAGE = flag.String("ri", "redis:7.0-alpine", "redis image name")
	APP_IMAGE = flag.String("ai", "walenpiotr/url-shortener:1.1.3", "app image name")
	CLIENT_IMAGE = flag.String("ci", "nginx:1.23", "client image name")
	NUM_SHARDS = flag.Int("n", 3, "number of redis shards")
	flag.Parse()

	if *NUM_SHARDS < 1 {
		log.Fatal("Invalid number of shards.")
	}
}

type AppConfig struct {
	Name string
	Port int
}

type RedisConfig struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type ClientConfig struct {
	Name string
	Port int
}

func createClientServiceConfig(clientConfig ClientConfig, appConfig AppConfig) types.ServiceConfig {
	dependsOn := types.DependsOnConfig{}
	dependsOn[appConfig.Name] = types.ServiceDependency{
		Condition: "service_started",
	}
	config := types.ServiceConfig{
		Name: clientConfig.Name,
		Ports: []types.ServicePortConfig{
			{
				Target:    uint32(clientConfig.Port),
				Published: strconv.Itoa(clientConfig.Port),
			},
		},
		DependsOn: dependsOn,
	}

	if *DEV_BUILD {
		config.Build = &types.BuildConfig{
			Context:    "./client",
			Dockerfile: "Dockerfile.prod",
		}
	} else if *DEV_HMR {
		config.Build = &types.BuildConfig{
			Context:    "./client",
			Dockerfile: "Dockerfile.dev",
		}
		config.Volumes = []types.ServiceVolumeConfig{
			{Type: "bind", Source: "./client", Target: "/node/app", Consistency: "delegated"}, // delegated = improved performance when changing files on host
			{Type: "volume", Target: "/node/app/node_modules"},                                // Anonymous volume to exclude hosts node_modules
		}
	} else {
		config.Image = *CLIENT_IMAGE
	}

	return config
}

func createAppServiceConfig(appConfig AppConfig, redisConfigs []RedisConfig) types.ServiceConfig {
	b, err := json.Marshal(redisConfigs)
	if err != nil {
		log.Fatal(err)
	}

	dependsOn := types.DependsOnConfig{}
	defaultServiceDependency := types.ServiceDependency{
		Condition: "service_started",
	}
	for _, c := range redisConfigs {
		dependsOn[c.Name] = defaultServiceDependency
	}

	config := types.ServiceConfig{
		Name: appConfig.Name,
		Environment: types.NewMappingWithEquals([]string{
			"REDIS_CONFIG=" + string(b),
		}),
		Ports: []types.ServicePortConfig{
			{
				Target:    uint32(appConfig.Port),
				Published: strconv.Itoa(appConfig.Port),
			},
		},
		DependsOn: dependsOn,
	}

	if *DEV_BUILD {
		config.Build = &types.BuildConfig{
			Context:    "./app",
			Dockerfile: "Dockerfile.prod",
		}
	} else if *DEV_HMR {
		config.Build = &types.BuildConfig{
			Context:    "./app",
			Dockerfile: "Dockerfile.dev",
		}
		config.Volumes = []types.ServiceVolumeConfig{
			{Type: "bind", Source: "./app", Target: "/app", Consistency: "delegated"}, // delegated = improved performance when changing files on host
		}
	} else {
		config.Image = *APP_IMAGE
	}

	return config
}

func createRedisServiceConfig(config RedisConfig) types.ServiceConfig {
	return types.ServiceConfig{
		Name:    config.Name,
		Image:   *REDIS_IMAGE,
		Restart: "always",
		Command: strings.Split("redis-server --save 20 1 --loglevel warning --requirepass "+config.Password, " "),
	}
}

func generateRedisConfigs() []RedisConfig {
	redisConfigs := []RedisConfig{}
	for i := 0; i < *NUM_SHARDS; i++ {
		iStr := strconv.Itoa(i)
		redisConfigs = append(
			redisConfigs,
			RedisConfig{Name: "redis-storage-" + iStr, Port: 6379, Password: "redis-storage-" + iStr},
		)
	}
	return redisConfigs
}

func main() {
	parseFlags()

	appConfig := AppConfig{Name: "go-app", Port: 8000}
	clientConfig := ClientConfig{Name: "client", Port: 5173}

	redisConfigs := generateRedisConfigs()

	services := types.Services{}
	for _, config := range redisConfigs {
		services = append(services, createRedisServiceConfig(config))
	}
	services = append(services, createAppServiceConfig(appConfig, redisConfigs))
	services = append(services, createClientServiceConfig(clientConfig, appConfig))

	project := types.Project{
		Services: services,
	}

	b, err := yaml.Marshal(project)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*FILENAME, b, 0644)
	log.Println("Generated " + *FILENAME + " file.")

	if err != nil {
		log.Fatal(err)
	}
}
