package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

type OutputType enumflag.Flag

const (
	Markdown OutputType = iota
)

var OutputTypeIds = map[OutputType][]string{
	Markdown: {"markdown"},
}

var schemaPath string
var targetPath string
var outputType OutputType

var rootCmd = &cobra.Command{
	Use:   "entmaid",
	Short: "A CLI for generating a mermaid.js Entity Relationship (ER) diagram for an Ent Schema, without needing a live database!",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := GenerateDiagram(schemaPath, targetPath, outputType); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&schemaPath, "schema", "s", "./ent/schema", "directory containing the schemas")
	rootCmd.PersistentFlags().StringVarP(&targetPath, "target", "t", "./ent/erd.md", "target directory for schemas")
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&outputType, "outputType", OutputTypeIds, enumflag.EnumCaseSensitive),
		"outputType", "o",
		"set the desired output type: can be 'markdown' (useful for GitHub)")
	rootCmd.PersistentFlags().Lookup("outputType").NoOptDefVal = "markdown"
}
