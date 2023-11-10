/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbcat",
	Short: "Display a Jupyter notebook in the terminal.",
	Long: `
'nbcat' is a command that displays a formatted version of the Jupyter
notebook specified. It is similar to the 'cat', 'bat', and 'glow' commands, but for 
a use case that none of them covers.`,

	Run: func(command *cobra.Command, args []string) {
		// var cell = FrameStrings({"testing here"}, 60)
		// fmt.Println(cell)

		var notebook = ParseNotebookFile("_scratch/demo.ipynb")

		var width = 100
		var representation = StringifyRawNotebook(notebook, width)

		fmt.Println(representation)
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
