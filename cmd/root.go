package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configFile string

// Help information
var rootCmd = &cobra.Command{
	Use:   "puck",
	Short: "A simple tool to check for apt package updates on multiple servers.",
	Long: `Puck (Package Update Checking Kit) is a simple tool that connects to your servers and checks for
apt package updates.`,
}

// Executes cobra
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Adds flags
func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "puck.yml", "puck configuration file")
}