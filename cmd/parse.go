package cmd

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

type BundleConfig struct {
}

type BundleLintConfig struct {
	notifications_in_prod bool
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

func ParseBundleLintConfig(path string) (*BundleLintConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read package config: %w", err)
	}

	var config map[string]interface{}
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse package config: %w", err)
	}

	tool, ok := config["tool"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("[tool] section not found in package config")
	}

	bundlelint, ok := tool["bundlelint"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("[tool.bundlelint] section not found in package config")
	}

	var lintConfig BundleLintConfig
	if err := mapToStruct(bundlelint, &lintConfig); err != nil {
		return nil, fmt.Errorf("failed to parse [tool.bundlelint] section: %w", err)
	}

	return &lintConfig, nil
}

func mapToStruct(input map[string]interface{}, output interface{}) error {
	serialized, err := yaml.Marshal(input)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(serialized, output)
}
