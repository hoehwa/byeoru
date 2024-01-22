/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hoehwa/byeoru/internal"
	"github.com/spf13/cobra"
)

// dealWithGenCmd represents the dealWithGen command.
var dealWithGenCmd = &cobra.Command{
	Use:   "dealWithGen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrettyPrint("/refactor/dealWithGen")
	},
}

func init() {
	refactorCmd.AddCommand(dealWithGenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dealWithGenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dealWithGenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
