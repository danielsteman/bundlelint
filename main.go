package main

import (
	"fmt"
	"github.com/danielsteman/bundlelint/cmd"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "bundlelint [bundle_path]",
	Short:   "A CLI to govern your Databricks asset bundles with flexibility",
	Long:    `A CLI to govern your Databricks asset bundles with flexibility. When the number of asset bundles in your company grows, you might want to set some rules. databricks-cli just validates the asset bundle config, but BundleLint checks if best practices are applied.`,
	Version: version,
	Args:    cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		bundleDir := "."
		configFile := "./pyproject.toml"

		if len(args) > 0 {
			if filepath.IsAbs(args[0]) {
				bundleDir = args[0]
			} else {
				cwd, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(c.OutOrStderr(), "Failed to get current working directory: %v\n", err)
					return
				}
				bundleDir = filepath.Join(cwd, args[0])
				configFile = filepath.Join(bundleDir, "pyproject.toml")
			}
		}

		fileInfo, err := os.Stat(bundleDir)
		if os.IsNotExist(err) || !fileInfo.IsDir() {
			fmt.Fprintf(c.OutOrStderr(), "Error: Bundle configuration directory not found: %s\n", bundleDir)
			return
		}

		bundleConfigPath := filepath.Join(bundleDir, "databricks.yml")

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			fmt.Fprintf(c.OutOrStderr(), "Error: Lint config file not found: %s\n", configFile)
			return
		}

		fmt.Fprintf(c.OutOrStdout(), "Validating bundle configuration: %s\n", bundleConfigPath)

		bundleConfig, err := cmd.ParseBundleConfig(bundleConfigPath)
		if err != nil {
			fmt.Fprintf(c.OutOrStderr(), "Error parsing bundle config: %s\n", err)
			return
		}

		lintConfig, err := cmd.ParseLintConfig(configFile)
		if err != nil {
			fmt.Fprintf(c.OutOrStderr(), "Error parsing lint config: %s\n", err)
			return
		}

		if cmd.ValidateConfigs(bundleConfig, lintConfig) {
			fmt.Fprintf(c.OutOrStdout(), "Validation successful!\n")
		} else {
			fmt.Fprintf(c.OutOrStdout(), "Validation failed!\n")
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
