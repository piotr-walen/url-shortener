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
var DEV *bool
var REDIS_IMAGE *string
var APP_IMAGE *string
var NUM_SHARDS *int

func parseFlags() {
	FILENAME = flag.String("f", "docker-compose.yml", "a filename of generated docker compose file")
	DEV = flag.Bool("d", false, "generate dev file")
	REDIS_IMAGE = flag.String("ri", "redis:7.0-alpine", "redis image name")
	APP_IMAGE = flag.String("ai", "walenpiotr/url-shortener:1.1.3", "app image name")
	NUM_SHARDS = flag.Int("n", 3, "number of redis shards")
	flag.Parse()

	if *NUM_SHARDS < 1 || *NUM_SHARDS > 10 {
		log.Fatal("Number of redis shards should be between 1 and 10")
	}
}

type AppConfig struct {
	Name string
	Port int
}

type RedisConfig struct {
	Name         string `json:"name"`
	Port         int    `json:"port"`
	ExternalPort string `json:"-"`
	Password     string `json:"password"`
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

	if *DEV {
		config.Build = &types.BuildConfig{
			Context: ".",
		}
	} else {
		config.Image = *APP_IMAGE
	}

	return config
}

func createRedisServiceConfig(config RedisConfig) types.ServiceConfig {
	return types.ServiceConfig{
		Name:  config.Name,
		Image: *REDIS_IMAGE,
		Ports: []types.ServicePortConfig{
			{
				Target:    uint32(config.Port),
				Published: config.ExternalPort,
			},
		},
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
			RedisConfig{Name: "redis-storage-" + iStr, Port: 6379, ExternalPort: "6379" + iStr, Password: "redis-storage-" + iStr},
		)
	}
	return redisConfigs
}

func main() {
	parseFlags()

	appConfig := AppConfig{Name: "go-app", Port: 8000}
	redisConfigs := generateRedisConfigs()

	services := types.Services{}
	for _, config := range redisConfigs {
		services = append(services, createRedisServiceConfig(config))
	}
	services = append(services, createAppServiceConfig(appConfig, redisConfigs))

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
