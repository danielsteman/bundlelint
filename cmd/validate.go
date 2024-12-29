package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate a Databricks asset bundle config",
	Long:  "Validate a Databricks asset bundle configuration file against user-defined rules.",
	Args:  cobra.MaximumNArgs(1), // Optional argument
	Run: func(cmd *cobra.Command, args []string) {
		// Default to "pyproject.toml" if no argument is provided
		file := "pyproject.toml"
		if len(args) > 0 {
			file = args[0]
		}

		// Check if the file exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("❌ Configuration file not found: %s\n", file)
			return
		}

		fmt.Printf("✅ Validating configuration file: %s\n", file)

		// Add your validation logic here
		// Example: validateConfig(file)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
