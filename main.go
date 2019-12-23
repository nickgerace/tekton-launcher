/*
TEKTON LAUNCHER
created by: Nick Gerace

MIT License, Copyright (c) Nick Gerace
See "LICENSE" file for more information

Please find license and further
information via the link below.
https://github.com/nickgerace/tekton-launcher
*/

package main

import (
	"fmt"
	"github.com/nickgerace/tekton-launcher/util"
	"github.com/spf13/cobra"
	"os"
)

// Add the main invocation command.
var rootCmd = &cobra.Command{
	Use:   "tekton-launcher",
	Short: "A brief description of the application",
	Long:  `A longer description that spans multiple lines for the application`,
}

// Add the "run" subcommand.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of the command",
	Long:  `A longer description that spans multiple lines for the command`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Launch()
	},
}

// Add the "version" subcommand.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of the command",
	Long:  `A longer description that spans multiple lines for the command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 0.1.0")
	},
}

// Initialize the subcommands.
func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute subcommands as part of main invocation function.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
