package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestParseBundleConfig(t *testing.T) {
	config, err := ParseBundleConfig("../test_bundle/databricks.yml")
	if err != nil {
		t.Fatalf("Failed to parse bundle config: %v", err)
	}
	prettyConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal config: %v", err)
	}

	fmt.Println("Config:", string(prettyConfig))
}

func TestParseLintConfig(t *testing.T) {
	config, err := ParseLintConfig("../test_bundle/pyproject.toml")
	if err != nil {
		t.Fatalf("Failed to parse bundle config: %v", err)
	}
    if config.NotificationsInProd != true {
		t.Fatalf("Expected value to be true: %v", err)
    }
	prettyConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal config: %v", err)
	}

	fmt.Println("Config:", string(prettyConfig))
}
