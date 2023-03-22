package config

import (
	"errors"
	"fmt"
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
	Access   string
	Secret   string
	Region   string
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
	Aws      AwsConfig
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

	err = loadEnvVars()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func loadEnvVars() error {
	access := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")
	endpoint := os.Getenv("AWS_ENDPOINT")

	if access == "" && config.Aws.Access == "" {
		return fmt.Errorf("either AWS_ACCESS_KEY_ID or aws.access is required")
	}

	if secret == "" && config.Aws.Secret == "" {
		return fmt.Errorf("either AWS_SECRET_ACCESS_KEY or aws.secret is required")
	}

	if region == "" && config.Aws.Region == "" {
		return fmt.Errorf("either AWS_DEFAULT_REGION or aws.region is required")
	}

	if access != "" {
		config.Aws.Access = access
	}

	if secret != "" {
		config.Aws.Secret = secret
	}

	if region != "" {
		config.Aws.Region = region
	}

	if endpoint != "" {
		config.Aws.Endpoint = endpoint
	}

	return nil
}
