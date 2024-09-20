package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config defines the format of the config we expect
type Config struct {
	Name            string   `yaml:"name"`
	Version         string   `yaml:"version"`
	ExpectedOutput  string   `yaml:"expected_output"`
	ScriptArguments []string `yaml:"arguments"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadConfigFromEnv(string(file))
}

func LoadConfigFromEnv(config string) (*Config, error) {
	replaced := os.ExpandEnv(config)
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(replaced), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
