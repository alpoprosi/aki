package config

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

const defaultPort = "8080"

type Config struct {
	Addr string `env:"HTTP_ADDR" envDefault:"127.0.0.1"`
	Port string `env:"HTTP_PORT" envDefault:"8888"`

	PgDSN string `env:"PG_DSN,required"`

	YAMLPath string `env:"CONFIG_YAML" envDefault:"config.yaml"`

	ApiVer       int      `yaml:"api_version"`
	AllowOrigins []string `yaml:"allow_origins"`
	LoginURL     string   `yaml:"login_url"`
}

func New() (*Config, error) {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		return nil, fmt.Errorf("parsing envs: %w", err)
	}

	if err := parseYAML(conf); err != nil {
		return nil, fmt.Errorf("parsing yaml: %w", err)
	}

	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return conf, nil
}

func parseYAML(c *Config) error {
	if c.YAMLPath == "" {
		return nil
	}

	f, err := os.Open(c.YAMLPath)
	if err != nil {
		return fmt.Errorf("opening config: %w", err)
	}
	defer f.Close()

	confBytes, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	if err = yaml.Unmarshal(confBytes, c); err != nil {
		return fmt.Errorf("unmarshall config: %w", err)
	}

	return nil
}

func (c *Config) validate() error {
	if ip := net.ParseIP(c.Addr); ip == nil && c.Addr != "localhost" {
		return fmt.Errorf("invalid http address: %s", c.Addr)
	}

	p, err := strconv.Atoi(c.Port)
	if err != nil {
		return fmt.Errorf("incorrect port value: %w", err)
	}

	if p < 0 {
		return fmt.Errorf("invalid http port: %s", c.Port)
	}

	if p == 0 {
		c.Port = defaultPort
	}

	if c.ApiVer < 1 {
		return fmt.Errorf("invalid api version")
	}

	return nil
}

func (c *Config) HTTPAddr() string {
	return net.JoinHostPort(c.Addr, c.Port)
}

func (c *Config) PrivatePrefix() string {
	return fmt.Sprintf("api/v%d", c.ApiVer)
}

func (c *Config) PublicPrefix() string {
	return fmt.Sprintf("pub/v%d", c.ApiVer)
}
