package cmd

import (
	"fmt"
	"os"

	// "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

type BundleConfig struct {
}

type BundleLintConfig struct {
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

	// _, err = jobs.Job(config)
	//
	// if err != nil {
	// 	return nil, fmt.Errorf("bundle config validation failed: %w", err)
	// }

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

func ValidateConfigs(bundleConfig *BundleConfig, bundleLintConfig *BundleLintConfig) bool {
	return true
}

var validateCmd = &cobra.Command{
	Use:   "validate [file]",
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

		bundleLintConfig, err := ParseBundleLintConfig(configFile)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error parsing package config: %s\n", err)
			return
		}

		if ValidateConfigs(bundleConfig, bundleLintConfig) {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation successful!\n")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Validation failed!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
