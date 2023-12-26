/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package ptrn

import (
	"fmt"

	"github.com/hoehwa/byeoru/cmd"
	"github.com/spf13/cobra"
)

// ptrnCmd represents the ptrn command.
var ptrnCmd = &cobra.Command{
	Use:   "ptrn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ptrn called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(ptrnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptrnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptrnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
