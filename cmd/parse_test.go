package cmd

import (
	"testing"
)

func TestParseBundleConfig(t *testing.T) {
	_, err := ParseBundleConfig("../test_bundle/databricks.yml")
	if err != nil {
		t.Fatalf("Failed to parse bundle config: %v", err)
	}
}
