/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/hoehwa/byeoru/cmd"
	_ "github.com/hoehwa/byeoru/cmd/ptrn"
	_ "github.com/hoehwa/byeoru/cmd/ptrn/design"
	_ "github.com/hoehwa/byeoru/cmd/ptrn/perf"
	_ "github.com/hoehwa/byeoru/cmd/ptrn/render"
	_ "github.com/hoehwa/byeoru/cmd/refactor"
	_ "github.com/hoehwa/byeoru/cmd/smell"
)

func main() {
	cmd.Execute()
}
