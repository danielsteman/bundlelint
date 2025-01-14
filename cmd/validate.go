package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func ValidateConfigs(bundleConfig *BundleConfig, lintConfig *LintConfig) bool {

	return true
}

var validateCmd = &cobra.Command{
	Use:   "lint [bundle_path]",
	Short: "Lint a Databricks asset bundle config",
	Long:  "Lint a Databricks asset bundle configuration file against user-defined rules.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bundleDir := "."
		configFile := "./pyproject.toml"

		if len(args) > 0 {
			if filepath.IsAbs(args[0]) {
				bundleDir = args[0]
			} else {
				cwd, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "Failed to get current working directory: %v\n", err)
					return
				}
				bundleDir = filepath.Join(cwd, args[0])
			}
		}

		fileInfo, err := os.Stat(bundleDir)
		if os.IsNotExist(err) || !fileInfo.IsDir() {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: Bundle configuration directory not found: %s\n", bundleDir)
			return
		}

		bundleConfigPath := filepath.Join(bundleDir, "databricks.yml")

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: Lint config file not found: %s\n", configFile)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Validating bundle configuration: %s\n", bundleDir)

		bundleConfig, err := ParseBundleConfig(bundleConfigPath)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error parsing bundle config: %s\n", err)
			return
		}

		lintConfig, err := ParseLintConfig(configFile)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error parsing lint config: %s\n", err)
			return
		}

		if ValidateConfigs(bundleConfig, lintConfig) {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation successful!\n")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation failed!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
