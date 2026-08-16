/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"godrop/internal/server"
	"log"

	"github.com/spf13/cobra"
)

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Println(err)
		}
		server.StartServerDownload(path)
	},
}

func init() {
	dropCmd.Flags().String("path", "", "Set path of file")
	rootCmd.AddCommand(dropCmd)
}
