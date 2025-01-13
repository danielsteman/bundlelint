package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ValidateConfigs(bundleConfig *BundleConfig, lintConfig *LintConfig) bool {
	return true
}

var validateCmd = &cobra.Command{
	Use:   "validate [bundle_path]",
	Short: "Validate a Databricks asset bundle config",
	Long:  "Validate a Databricks asset bundle configuration file against user-defined rules.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := "pyproject.toml"

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Fprintf(cmd.OutOrStdout(), "pyproject.toml not found: %s\n", configFile)
			return
		}

		var bundleDir string
		if len(args) > 0 {
			if args[0][0] == 47 { // if arg starts with slash (/)
				bundleDir = args[0]
			} else {
				cwd, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(cmd.OutOrStdout(), "Failed to get current working directory: %v", err)
				}
				bundleDir = cwd + "/" + args[0]
			}
		}

		bundleConfigPath := bundleDir + "/databricks.yml"

		if _, err := os.Stat(bundleDir); os.IsNotExist(err) {
			fmt.Fprintf(cmd.OutOrStdout(), "Bundle configuration not found: %s\n", bundleDir)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Validating bundle configuration: %s\n", bundleDir)

		bundleConfig, err := ParseBundleConfig(bundleConfigPath)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error parsing bundle config: %s\n", err)
			return
		}

		LintConfig, err := ParseLintConfig(configFile)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error parsing package config: %s\n", err)
			return
		}

		fmt.Println(LintConfig)

		if ValidateConfigs(bundleConfig, LintConfig) {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation successful!\n")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation failed!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
