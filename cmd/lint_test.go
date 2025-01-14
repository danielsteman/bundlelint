package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
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
	if err := os.Mkdir(bundleDir, 0755); err != nil {
		t.Fatalf("Failed to create temporary bundle directory: %v", err)
	}
	defer os.RemoveAll(bundleDir)

	bundleConfig := []byte(`
bundle:
  name: "Test Bundle"
include: []
targets: {}
resources:
  jobs: {}
`)
	bundleConfigPath := bundleDir + "/databricks.yml"
	if err := os.WriteFile(bundleConfigPath, bundleConfig, 0644); err != nil {
		t.Fatalf("Failed to create temporary bundle config file: %v", err)
	}

	testRootCmd := &cobra.Command{
		Use: "bundlelint",
		Run: rootCmd.Run,
	}

	output := executeCommand(testRootCmd, bundleDir)

	expected := "Validating bundle configuration: test_bundle\nValidation successful!\n"
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
}

func executeCommand(root *cobra.Command, args ...string) string {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_ = root.Execute()
	return buf.String()
}
