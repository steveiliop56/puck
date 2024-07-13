package updatechecker

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/ssh"
	"github.com/steveiliop56/puck/internal/validators"
)

func UpdateCache(server config.Server) (string, error) {
	var validaterErr = validators.ValidateServer(server)
	if validaterErr != nil {
		return "", validaterErr
	}

	var command = "echo " + server.Password + "| sudo -S apt update"
	var sshOutput, sshErr = ssh.RunCommandRich(server, command)
	if sshErr != nil {
		return sshOutput, sshErr
	}

	return sshOutput, nil
}

func GetUpgradable(server config.Server) (bool, string, error) {
	var validaterErr = validators.ValidateServer(server)
	if validaterErr != nil {
		return false, "", validaterErr
	}

	var command = "echo " + server.Password + "| sudo -S apt list --upgradable"
	var sshOutput, sshErr = ssh.RunCommandRich(server, command)
	if sshErr != nil {
		return false, sshOutput, sshErr
	}

	if sshOutput == "Listing... Done" {
		return true, sshOutput, nil
	}

	return false, sshOutput, nil
}

func GetUpgrades(server config.Server) {
	var _, updateCacheErr = UpdateCache(server)
	if updateCacheErr != nil {
		color.Set(color.FgRed)
		fmt.Print("\n✗ ")
		color.Unset()
		fmt.Printf("Failed to update cache of server %s, error: %s\n", server.Name, updateCacheErr)
		return;
	}
	var hasUpdate, _, upgradableErr = GetUpgradable(server)
	if upgradableErr != nil {
		color.Set(color.FgRed)
		fmt.Print("\n✗ ")
		color.Unset()
		fmt.Printf("Failed to get upgradable packages of server %s, error: %s\n", server.Name, upgradableErr)
		return;
	}
	if hasUpdate {
		color.Set(color.FgGreen)
		fmt.Print("\n✔ ")
		color.Unset()
		fmt.Printf("Server %s has available updates!\n", server.Name)
	}
	fmt.Printf("\nNo updates for server %s\n", server.Name)
}