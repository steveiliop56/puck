package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/puck/internal/constants"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// Just prints the version from the constants
var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Get the cli version.",
	Long: "This command simply prints the current cli version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Puck CLI version: ")
		color.Set(color.FgGreen)
		fmt.Printf("%s\n", constants.Version)
		color.Unset()
	},
}