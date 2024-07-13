package updatechecker

import (
	"strings"

	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/ssh"
	"github.com/steveiliop56/puck/internal/utils"
	"github.com/steveiliop56/puck/internal/validators"
)

// Given a server struct and the package manager command it runs the command and returns the output or the error
func UpdateCache(server config.Server, pacakgeManagerCommand string) (string, error) {
	var validaterErr = validators.ValidateServer(server)
	if validaterErr != nil {
		return "", validaterErr
	}

	var command = ""

	if server.NoSudo {
		command = pacakgeManagerCommand
	} else {
		command = "echo " + server.Password + "| sudo -S " + pacakgeManagerCommand
	}

	var sshOutput, sshErr = ssh.RunCommandRich(server, command)
	if sshErr != nil {
		return sshOutput, sshErr
	}

	return sshOutput, nil
}

// Given a server struct and the package manager command it runs the command and returns the output and if the system has upgrades or the error
func GetUpgradable(server config.Server, pacakgeManagerCommand string) (bool, string, error) {
	var validaterErr = validators.ValidateServer(server)
	if validaterErr != nil {
		return false, "", validaterErr
	}

	var command = "";

	if server.NoSudo {
		command = pacakgeManagerCommand
	} else {
		command = "echo " + server.Password + "| sudo -S " + pacakgeManagerCommand
	}

	var sshOutput, sshErr = ssh.RunCommandRich(server, command)
	if sshErr != nil {
		return false, sshOutput, sshErr
	}

	if strings.Trim(sshOutput, "\n") == "0" {
		return false, sshOutput, nil
	}

	return true, sshOutput, nil
}

// Given a server struct it checks for updates/upgrades and returns if the server has updates, if the server is skipepd and if there was an error
func GetUpgrades(server config.Server) (bool, bool, error) {
	var distro, distroErr = utils.GetDistro(server)
	if distroErr != nil {
		return false, false, distroErr
	}
	var pacakgeManagerCommand, skipped = utils.GetCommand(distro)
	if skipped {
		return false, true, nil
	}
	var _, updateCacheErr = UpdateCache(server, pacakgeManagerCommand[0])
	if updateCacheErr != nil {
		return false, false, updateCacheErr;
	}
	var hasUpdate, _, upgradableErr = GetUpgradable(server, pacakgeManagerCommand[1])
	if upgradableErr != nil {
		return false, false, upgradableErr;
	}
	if hasUpdate {
		return true, false, nil;
	}
	return false, false, nil;
}