package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Generate code for flows",
	Long: `Generate (a) code file(s) for every flow file.
The code is generated according to the given template file.
The template file is generating the file names of the generated code files, too.
The content of the 'outputDir' flag is given to the file name template.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("code called")
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
