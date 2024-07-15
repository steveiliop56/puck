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
	"github.com/steveiliop56/puck/internal/notifications"
	"github.com/steveiliop56/puck/internal/updatechecker"
	"github.com/steveiliop56/puck/internal/validators"
)

// Define flag variable
var doNtfy bool;

// Adds command to cobra
func init() {
	checkCmd.Flags().BoolVarP(&doNtfy, "ntfy", "n", false, "Send a notification via ntfy when check is done.")
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
			fmt.Print("✗ ")
			color.Unset()
			fmt.Println(" Cannot read config!")
			os.Exit(1)
		}

		// Parses the config file
		var config config.Config
		var unmarshalErr = yaml.Unmarshal(configRaw, &config)
		if unmarshalErr != nil {
			color.Set(color.FgRed)
			fmt.Print("✗ ")
			color.Unset()
			fmt.Println("Cannot parse config!")
			os.Exit(1)
		}

		// Validates config file
		var validateErr = validators.ValidateConfig(config)
		if validateErr != nil {
			color.Set(color.FgRed)
			fmt.Print("✗ ")
			color.Unset()
			fmt.Println("Cannot validate config!")
			os.Exit(1)
		}

		// Verify ntfy url exists if ntfy is enabled
		if doNtfy && config.NtfyURL == "" {
			color.Set(color.FgYellow)
			fmt.Print("● ")
			color.Unset()
			fmt.Println("Ntfy notifications enabled but ntfy url not set, notification won't be sent!")
			doNtfy = false;
		}

		// Create 3 lists for skipped, upgradable and up to date servers
		var upgradable = []string{}
		var skipped = []string{}
		var uptodate = []string{}

		// Create spinner
		var spinnerAnimation = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		var spinner = spinner.New(spinnerAnimation, 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Checking for updates...\n"
		spinner.Start()

		// For each server check if its upgradable and add it in the upgradable list, if its skipped add it to the skipped list, if its up to date add it to the uptodate list and if there is an error stop the cli
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
			uptodate = append(uptodate, element.Name)
		}

		// Stops the spinner and shows the check finished message
		spinner.Stop()

		color.Set(color.FgGreen)
		fmt.Print("✔ ")
		color.Unset()
		fmt.Println("Update check finished!")
		
		// Prints each server's status (skipped, upgradable, uptodate)
		for _, element := range upgradable {
			color.Set(color.FgBlue)
			fmt.Print("↻ ")
			color.Unset()
			fmt.Printf("Server %s has an update!\n", element)
		}

		for _, element := range skipped {
			color.Set(color.FgYellow)
			fmt.Print("● ")
			color.Unset()
			fmt.Printf("Server %s skipped, unsupported distro.\n", element)
		}

		for _, element := range uptodate {
			color.Set(color.FgGreen)
			fmt.Print("✔ ")
			color.Unset()
			fmt.Printf("Server %s is up to date!\n", element)
		}

		// Check if ntfy is enabled and if it is send a notification
		if doNtfy {
			notifications.NotifyNtfy(skipped, upgradable, uptodate, config.NtfyURL)
		}
	},
}