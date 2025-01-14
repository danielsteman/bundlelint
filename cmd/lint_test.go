package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestValidateCommand_DefaultFile(t *testing.T) {
	tempFile := "pyproject.toml"
	content := []byte(`[tool.bundlelint]
notifications_in_prod = true`)
	if err := os.WriteFile(tempFile, content, 0644); err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile)

	bundleDir := "test_bundle"
	if err := os.MkdirAll(bundleDir, 0755); err != nil {
		t.Fatalf("Failed to create temporary bundle directory: %v", err)
	}
	defer os.RemoveAll(bundleDir)

	bundleConfig := []byte(`
bundle:
  name: "test_bundle"
include: []
targets:
  dev:
    mode: "development"
    default: true
    workspace:
      host: "https://workspace-id.cloud.databricks.com"
resources: {}
`)
	bundleConfigPath := bundleDir + "/databricks.yml"
	if err := os.WriteFile(bundleConfigPath, bundleConfig, 0644); err != nil {
		t.Fatalf("Failed to create temporary bundle config file: %v", err)
	}

	output := new(bytes.Buffer)
	rootCmd := NewRootCmd()
	rootCmd.SetOut(output)
	rootCmd.SetErr(output)
	rootCmd.SetArgs([]string{bundleDir})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	outputStr := output.String()
	expectedParts := []string{
		"Validating bundle configuration:",
		"Validation successful!",
	}

	for _, part := range expectedParts {
		if !strings.Contains(outputStr, part) {
			t.Errorf("Expected output to contain %q but it was not found in %q", part, outputStr)
		}
	}
}
