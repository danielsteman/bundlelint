package cmd

import (
	"fmt"
	"testing"
)

func TestParseBundleConfig(t *testing.T) {
	config, err := ParseBundleConfig("../test_bundle/databricks.yml")
	if err != nil {
		t.Fatalf("Failed to parse bundle config: %v", err)
	}
	fmt.Println(config)
}
