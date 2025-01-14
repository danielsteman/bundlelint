package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "bundlelint validate",
	Short:   "A CLI to govern your Databricks asset bundles with flexibility",
	Long:    `A CLI to govern your Databricks asset bundles with flexibility. When the number of asset bundles in your company grows, you might want to set some rules. databricks-cli just validates the asset bundle config, but BundleLint checks if best practices are applied.`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("BundleLint CLI v%s\n", version)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
