package cmd

import (
	"fmt"
	"os"

	"github.com/mcpvault/mcpvault/internal/config"
	"github.com/spf13/cobra"
)

var (
	removeForce bool
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove NAME",
	Aliases: []string{"rm"},
	Short:   "Remove a configuration",
	Long: `Remove a configuration from MCPVault.

Example:
  mcpv remove weather
  mcpv remove weather --force`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		manager, err := config.NewManager()
		if err != nil {
			return err
		}

		// Check if configuration exists
		_, err = manager.Get(name)
		if err != nil {
			return err
		}

		// Ask for confirmation unless --force flag is used
		if !removeForce {
			fmt.Printf("Are you sure you want to remove configuration '%s'? [y/N]: ", name)
			var response string
			fmt.Scanln(&response)
			if response != "y" && response != "Y" {
				fmt.Println("Operation cancelled")
				return nil
			}
		}

		if err := manager.Remove(name); err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "Configuration '%s' removed successfully\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Optional flags
	removeCmd.Flags().BoolVarP(&removeForce, "force", "f", false, "Force removal without confirmation")
} 