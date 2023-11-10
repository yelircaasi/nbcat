/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var placeholderCmd = &cobra.Command{
	Use:   "ph",
	Short: "This is just a placeholder.",
	Long: `This here is 
THE MULTILINE PLACEHOLDER!`,

	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(placeholderCmd)

	// persistent flag, which will work for this command and all subcommands
	placeholderCmd.PersistentFlags().String("foo", "", "help message for foo")

	// local flag which will only run when this command is called directly
	placeholderCmd.Flags().BoolP("localcommand", "l", false, "help message for localcommand")
}
