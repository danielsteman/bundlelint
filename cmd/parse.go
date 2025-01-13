package cmd

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

type BundleName struct {
	Name string `yaml:"name"`
}

type BundleConfig struct {
	Bundle BundleName `yaml:"bundle"`
}

type LintConfig struct {
	NotificationsInProd bool `toml:"notifications_in_prod"`
}

func ParseBundleConfig(path string) (*BundleConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read bundle config: %w", err)
	}

	var config BundleConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse bundle config: %w", err)
	}

	return &config, nil
}

func ParseLintConfig(path string) (*LintConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read package config: %w", err)
	}

	var config struct {
		Tool struct {
			BundleLint *LintConfig `toml:"bundlelint"`
		} `toml:"tool"`
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse package config: %w", err)
	}

	if config.Tool.BundleLint == nil {
		return nil, fmt.Errorf("[tool.bundlelint] section not found in package config")
	}

	return config.Tool.BundleLint, nil
}
