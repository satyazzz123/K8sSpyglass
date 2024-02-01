/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/satyazzz123/K8sSpyglass.git/cmd/check"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spyk8",
	Short: "A small application to spy on your k8s cluster :)",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(check.CheckCmd)
	// Hide the default "completion" sub-command
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}
