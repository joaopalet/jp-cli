/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "jp-cli/cmd"

var (
	version = "dev"
)

func main() {
	cmd.ExecuteRootCmd(version)
}
