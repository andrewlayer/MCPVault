package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/mcpvault/mcpvault/internal/config"
	"github.com/spf13/cobra"
)

var (
	listVerbose bool
)

// ANSI escape codes for text formatting
const (
	boldOn  = "\033[1m"
	boldOff = "\033[0m"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all configurations",
	Long: `List all configurations stored in MCPVault.
Use the --verbose flag to show additional details.

Examples:
  mcpv list
  mcpv list --verbose`,
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := config.NewManager()
		if err != nil {
			return err
		}

		configs, err := manager.List(listVerbose)
		if err != nil {
			return err
		}

		if len(configs) == 0 {
			fmt.Fprintln(os.Stdout, "No configurations found")
			return nil
		}

		// Sort configurations alphabetically
		sort.Strings(configs)

		for _, config := range configs {
			if listVerbose {
				name := config[:getNameEndIndex(config)]
				details := config[getNameEndIndex(config):]
				fmt.Fprintf(os.Stdout, "%s%s%s%s\n", boldOn, name, boldOff, details)
				fmt.Fprintln(os.Stdout)
			} else {
				fmt.Fprintf(os.Stdout, "%s%s%s\n", boldOn, config, boldOff)
			}
		}

		return nil
	},
}

// getNameEndIndex finds where the name ends in the verbose output string
func getNameEndIndex(config string) int {
	// In verbose mode, the format is "name\n  Description: ..."
	// so we find the first newline character
	for i, char := range config {
		if char == '\n' {
			return i
		}
	}
	return len(config) // If no newline is found, return the whole string
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Optional flags
	listCmd.Flags().BoolVarP(&listVerbose, "verbose", "v", false, "Show verbose output including descriptions and timestamps")
} 