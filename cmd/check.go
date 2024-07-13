package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/updatechecker"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use: "check",
	Short: "Check all servers for updates.",
	Long: "Check all servers defined in the configuration file for updates.",
	Run: func(cmd *cobra.Command, args []string) {
		var configRaw, readErr = os.ReadFile(configFile)
		if readErr != nil {
			color.Set(color.FgRed)
			fmt.Print("\n✗ ")
			color.Unset()
			fmt.Print(" Cannot read config!")
			os.Exit(1)
		}

		var config config.Config
		var unmarshalErr = yaml.Unmarshal(configRaw, &config)
		if unmarshalErr != nil {
			color.Set(color.FgRed)
			fmt.Print("\n✗ ")
			color.Unset()
			fmt.Print("Cannot parse config!")
			os.Exit(1)
		}
		var spinnerAnimation = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		var spinner = spinner.New(spinnerAnimation, 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Checking for updates..."
		spinner.Start()
		for _, element := range config.Servers {
			spinner.Suffix = " Checking for updates on server " + element.Name
			updatechecker.GetUpgrades(element)
		}
		spinner.Stop()
		color.Set(color.FgGreen)
		fmt.Print("✔ ")
		color.Unset()
		fmt.Printf("Proccess finished!\n")
	},
}