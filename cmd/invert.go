/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"img-processor/internal"
)

// invertCmd represents the invert command
var invertCmd = &cobra.Command{
	Use:   "invert [path/to/image]",
	Short: "invert the image at the given path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.InvertImage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(invertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// invertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// invertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
