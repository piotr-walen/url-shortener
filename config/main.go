package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type primitive interface {
	string | int | bool
}

func noop(v string) (string, error) {
	return v, nil
}

func loadVariable[T primitive](
	name string,
	required bool,
	parse func(string) (T, error),
	report func(error),
) (variable T, loaded bool) {
	value, loaded := os.LookupEnv(name)
	if !loaded && required {
		report(errors.New("ENV: " + name + " is required"))
		return
	}
	if !loaded && !required {
		return
	}
	variable, err := parse(value)
	if err != nil {
		report(errors.New("ENV: " + name + " has invalid value: " + value))
		return
	}
	return
}

type reporter struct {
	reportedErrors []error
}

func (r *reporter) report(err error) {
	r.reportedErrors = append(r.reportedErrors, err)
}

func (r *reporter) mergeErrors() error {
	var err error
	if len(r.reportedErrors) > 0 {
		errStrings := []string{}
		for _, err := range r.reportedErrors {
			errStrings = append(errStrings, err.Error())
		}
		err = errors.New(strings.Join(errStrings, ", "))
	}
	return err
}

type Config struct {
	MAX_URL_LENGTH int
	ADDR           string
	LOG_TRAFFIC    bool
}

const (
	DEFAULT_ADDR           = ":8000"
	DEFAULT_MAX_URL_LENGTH = 6
	DEFAULT_LOG_TRAFFIC    = true
)

var config = Config{
	ADDR:           DEFAULT_ADDR,
	MAX_URL_LENGTH: DEFAULT_MAX_URL_LENGTH,
	LOG_TRAFFIC:    DEFAULT_LOG_TRAFFIC,
}

func ParseConfig() error {
	r := reporter{
		reportedErrors: []error{},
	}
	if ADDR, loaded := loadVariable("ADDR", false, noop, r.report); loaded {
		config.ADDR = ADDR
	}
	if LOG_TRAFFIC, loaded := loadVariable("LOG_TRAFFIC", false, strconv.ParseBool, r.report); loaded {
		config.LOG_TRAFFIC = LOG_TRAFFIC
	}
	if MAX_URL_LENGTH, loaded := loadVariable("MAX_URL_LENGTH", false, strconv.Atoi, r.report); loaded {
		config.MAX_URL_LENGTH = MAX_URL_LENGTH
	}
	return r.mergeErrors()
}

func GetConfig() *Config {
	return &config
}
