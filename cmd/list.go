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
			fmt.Fprintln(os.Stdout, config)
			if listVerbose {
				fmt.Fprintln(os.Stdout) // Add a blank line between verbose entries
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Optional flags
	listCmd.Flags().BoolVarP(&listVerbose, "verbose", "v", false, "Show verbose output including descriptions and timestamps")
} 