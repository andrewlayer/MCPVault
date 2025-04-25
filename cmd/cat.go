package cmd

import (
	"os"

	"github.com/mcpvault/mcpvault/internal/config"
	"github.com/spf13/cobra"
)

var (
	catFormat string
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat NAME",
	Short: "Display a configuration",
	Long: `Display the contents of a configuration.
By default, the output is in JSON format.

Examples:
  mcpv cat weather
  mcpv cat weather --format=json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		manager, err := config.NewManager()
		if err != nil {
			return err
		}

		return manager.PrintConfig(os.Stdout, name, catFormat)
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Optional flags
	catCmd.Flags().StringVarP(&catFormat, "format", "f", "json", "Output format (json, yaml)")
} 