/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerNewUserCmd represents the registerNewUser command
var registerNewUserCmd = &cobra.Command{
	Use:   "registerNewUser",
	Short: "Register a new user to the quiz",
	Long: `Register a new user to quiz with full name and e-mail that must be unique`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	DisableFlagsInUseLine: true,
	Example: `start server:
	fasttrack_api registerNewUser <name> <email>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("registerNewUser called")
	},
}

func init() {
	rootCmd.AddCommand(registerNewUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerNewUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerNewUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
