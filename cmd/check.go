package cmd

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/updatechecker"
)

// Adds command to cobra
func init() {
	rootCmd.AddCommand(checkCmd)
}

// Check command
var checkCmd = &cobra.Command{
	Use: "check",
	Short: "Check all servers for updates.",
	Long: "Check all servers defined in the configuration file for updates.",
	Run: func(cmd *cobra.Command, args []string) {
		// Reads the config file
		var configRaw, readErr = os.ReadFile(configFile)
		if readErr != nil {
			color.Set(color.FgRed)
			fmt.Print("\n✗ ")
			color.Unset()
			fmt.Print(" Cannot read config!")
			os.Exit(1)
		}

		// Parses the config file
		var config config.Config
		var unmarshalErr = yaml.Unmarshal(configRaw, &config)
		if unmarshalErr != nil {
			color.Set(color.FgRed)
			fmt.Print("\n✗ ")
			color.Unset()
			fmt.Print("Cannot parse config!")
			os.Exit(1)
		}

		// Create 2 lists for skipped and upgradable servers
		var upgradable = []string{}
		var skipped = []string{}

		// Create spinner
		var spinnerAnimation = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		var spinner = spinner.New(spinnerAnimation, 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Checking for updates...\n"
		spinner.Start()

		// For each server check if its upgradable and add it in the upgradable list, if its skipped add it to the skipped list and if there is an error stop the cli
		for _, element := range config.Servers {
			var hasUpdate, isSkipped, err = updatechecker.GetUpgrades(element)
			if err != nil {
				spinner.Stop()
				color.Set(color.FgRed)
				fmt.Print("✗ ")
				color.Unset()
				fmt.Printf("Error in getting updates for server %s. Error: %s\n", element.Name, err)
				os.Exit(1)
			}
			if hasUpdate {
				upgradable = append(upgradable, element.Name)
			}
			if isSkipped {
				skipped = append(skipped, element.Name)
			}
		}

		// Stops the spinner and shows the check finished message
		spinner.Stop()

		color.Set(color.FgGreen)
		fmt.Print("✔ ")
		color.Unset()
		fmt.Println("Update check finished!")
		
		// Prints each server's status (skipped, upgradable)
		for _, element := range config.Servers {
			if slices.Contains(upgradable, element.Name) {
				color.Set(color.FgBlue)
				fmt.Print("↻ ")
				color.Unset()
				fmt.Printf("Server %s has an update!\n", element.Name)
			} else if slices.Contains(skipped, element.Name) {
				color.Set(color.FgYellow)
				fmt.Print("● ")
				color.Unset()
				fmt.Printf("Server %s skipped, unsupported distro.\n", element.Name)
			} else {
				color.Set(color.FgGreen)
				fmt.Print("✔ ")
				color.Unset()
				fmt.Printf("Server %s is up to date!\n", element.Name)
			}
		}
	},
}