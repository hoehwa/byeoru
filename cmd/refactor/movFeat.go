/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package refactor

import (
	"github.com/hoehwa/byeoru/utills"
	"github.com/spf13/cobra"
)

// movFeatCmd represents the movFeat command.
var movFeatCmd = &cobra.Command{
	Use:   "movFeat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/refactor/movFeat")
	},
}

func init() {
	refactorCmd.AddCommand(movFeatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// movFeatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movFeatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
