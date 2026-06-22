package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1)),
	Short: "Generate documentation for flows",
	Long: `Generate a MarkDown file for every Flow file.
The MarkDown contains SVG images of the flows in the Flow file
and some documentation.

Exmples:
  flow-dsl doc ./example.flow
  flow-dsl doc --outputDir=doc/org/flowdev/example src/org/flowdev/example/example.flow

The first example creates a MarkDown file './example.md' with the desired documentation.
The second example creates a MarkDown file 'doc/org/flowdev/example/example.md' as it might be idiomatic in a Java project.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("doc called with args: %q\n", args)
		fmt.Printf("outputDir: %q\n", outputDir)
		flags := cmd.Flags()
		tpl, err := flags.GetString("template")
		fmt.Printf("template: %q, err: %v\n", tpl, err)
	},
}

var outputDir string

func init() {
}

func init() {
	rootCmd.AddCommand(docCmd)

    docCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "directory for generated documentation files (default is directory of input files)")
    viper.BindPFlag("outputDir", docCmd.Flags().Lookup("outputDir"))

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
