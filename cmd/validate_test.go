package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestValidateCommand_DefaultFile(t *testing.T) {
	// Create a temporary pyproject.toml file for testing
	tempFile := "pyproject.toml"
	content := []byte(`[tool.bundlelint]`)
	if err := os.WriteFile(tempFile, content, 0644); err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile)

	// Initialize rootCmd with validateCmd
	testRootCmd := &cobra.Command{Use: "bundlelint"}
	testRootCmd.AddCommand(validateCmd)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	println(cwd)

	s := "/hoi"
	strings.Split(s, "")
	println(s[0])

	// Capture the output of the command
	output := executeCommand(testRootCmd, "validate", "../test_bundle")
	println(output)
	//
	// // Check the output
	// expected := "Validating configuration file: pyproject.toml\n"
	// if output == expected {
	// 	t.Errorf("Expected %q but got %q", expected, output)
	// }
}

// func TestValidateCommand_MissingFile(t *testing.T) {
// 	// Initialize rootCmd with validateCmd
// 	testRootCmd := &cobra.Command{Use: "bundlelint"}
// 	testRootCmd.AddCommand(validateCmd)
//
// 	// Capture the output of the command
// 	output := executeCommand(testRootCmd, "validate", "pyproject.toml")
//
// 	// Check the output
// 	expected := "Configuration file not found: pyproject.toml\n"
// 	if output != expected {
// 		t.Errorf("Expected %q but got %q", expected, output)
// 	}
// }

// Helper function to execute a Cobra command and capture the output
func executeCommand(root *cobra.Command, args ...string) string {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_ = root.Execute() // Ignore errors for testing output
	return buf.String()
}
