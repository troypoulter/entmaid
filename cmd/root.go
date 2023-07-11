package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var targetSchemaPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "entmaid",
	Short: "A CLI for generating a mermaid.js Entity Relationship (ER) diagram for an Ent Schema, without needing a live database!",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := GenerateDiagram(targetSchemaPath); err != nil {
			return err
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetSchemaPath, "target", "t", "./ent/schema", "target directory for schemas")
}
