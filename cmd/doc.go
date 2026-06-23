package cmd

import (
	"fmt"

	"github.com/flowdev/flow-dsl/doc"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("doc called with args: %q\n", args)
		fmt.Printf("outputDir: %q\n", outputDir)
		flags := cmd.Flags()
		tpl, err := flags.GetString("template")
		fmt.Printf("template file: %q, err: %v\n", tpl, err)

		return doc.DocumentFlows(
			args,
			viper.GetBool("doc.links"),
			viper.GetInt("doc.width"),
			viper.GetBool("doc.dark"),
			nil,
		)
	},
}

var outputDir string
var links bool
var width int
var dark bool

func init() {
	rootCmd.AddCommand(docCmd)

    docCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "directory for generated documentation files (default is directory of input files)")
    viper.BindPFlag("doc.outputDir", docCmd.Flags().Lookup("outputDir"))
    docCmd.Flags().BoolVarP(&links, "links", "l", true, "should the documentation contain links?")
    viper.BindPFlag("doc.links", docCmd.Flags().Lookup("links"))
    docCmd.Flags().IntVarP(&width, "width", "w", 1900, "maximum width of a diagram")
    viper.BindPFlag("doc.width", docCmd.Flags().Lookup("width"))
    docCmd.Flags().BoolVarP(&dark, "dark", "d", false, "draw diagrams in dark mode?")
    viper.BindPFlag("doc.dark", docCmd.Flags().Lookup("dark"))
}
