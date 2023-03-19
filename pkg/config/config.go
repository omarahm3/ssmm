package config

import (
	"errors"
	"os"

	"github.com/jinzhu/configor"
)

var (
	ErrConfigNotFound      = errors.New("config file not found")
	ErrConfigFileIsEmpty   = errors.New("config file is empty")
	ErrInvalidConfigFormat = errors.New("invalid config format")
)

var (
	config = Config{}
)

type AwsConfig struct {
	Access   string `required:"true"`
	Secret   string `required:"true"`
	Region   string `required:"true"`
	Endpoint string
}

type Variable struct {
	Key       string `required:"true"`
	Value     string `required:"true"`
	Secure    bool   `default:"false"`
	Overwrite bool   `default:"false"`
}

type Project struct {
	Name         string `required:"true"`
	Environments []struct {
		Name      string     `required:"true"`
		Variables []Variable `required:"true"`
	} `required:"true"`
}

type Config struct {
	Aws      AwsConfig `required:"true"`
	Projects []Project `required:"true"`
}

func Load(p string) (*Config, error) {
	_, err := os.Stat(p)
	if err != nil && os.IsNotExist(err) {
		return nil, ErrConfigNotFound
	}

	return load(p)
}

func load(p string) (*Config, error) {
	err := configor.Load(&config, p)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
