package cmd

import (

	"github.com/spf13/cobra"
)

// Version is set during build via ldflags
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "mcpv",
	Short: "MCPVault - A CLI tool for managing MCP server configurations",
	Long: `MCPVault (mcpv) is a CLI tool for managing MCP server configurations in a key-value 
store approach. It allows developers to add, remove, list, and view configurations 
for their MCP servers.`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
} 