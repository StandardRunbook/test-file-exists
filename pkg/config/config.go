package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// FileCheckConfig defines the format of the config we expect
type FileCheckConfig struct {
	Name            string   `yaml:"name"`
	Version         string   `yaml:"version"`
	ExpectedOutput  string   `yaml:"expected_output"`
	ScriptArguments []string `yaml:"arguments"`
}

func LoadConfigFromEnv(config string) (*FileCheckConfig, error) {
	replaced := os.ExpandEnv(config)
	cfg := &FileCheckConfig{}
	err := yaml.Unmarshal([]byte(replaced), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
