/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package refactor

import (
	"fmt"

	"github.com/hoehwa/byeoru/cmd"
	"github.com/spf13/cobra"
)

// refactorCmd represents the refactor command.
var refactorCmd = &cobra.Command{
	Use:   "refactor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		you can use following sub commands:
		- byeoru refactor cmpsMthd
		- byeoru refactor dealWithGen
		- byeoru refactor movFeat
		- byeoru refactor orgData
		- byeoru refactor simpCondexp
		- byeoru refactor simpMthdCalls
		`)
	},
}

func init() {
	cmd.RootCmd.AddCommand(refactorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refactorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refactorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
