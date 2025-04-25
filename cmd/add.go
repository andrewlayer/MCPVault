package cmd

import (
	"fmt"
	"os"

	"github.com/mcpvault/mcpvault/internal/config"
	"github.com/spf13/cobra"
)

var (
	addName        string
	addDescription string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [file path | json string]",
	Short: "Add a new configuration",
	Long: `Add a new configuration to MCPVault.
The configuration can be provided as a file path to a JSON file
or directly as a JSON string.

Examples:
  mcpv add config.json --name weather --description "Weather service config"
  mcpv add '{"api_key":"abc123","endpoint":"https://api.weather.com"}' --name weather`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if addName == "" {
			return fmt.Errorf("name is required")
		}

		manager, err := config.NewManager()
		if err != nil {
			return err
		}

		input := args[0]

		// Check if input is a file path
		if _, err := os.Stat(input); err == nil {
			return manager.AddFromFile(addName, addDescription, input)
		}

		// Assume input is a JSON string
		return manager.AddFromJSON(addName, addDescription, input)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Required flags
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Name of the configuration (required)")
	addCmd.MarkFlagRequired("name")

	// Optional flags
	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "Description of the configuration")
} 