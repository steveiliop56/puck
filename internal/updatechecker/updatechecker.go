package updatechecker

import (
	"strings"

	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/ssh"
	"github.com/steveiliop56/puck/internal/validators"
)

func UpdateCache(server config.Server) (string, error) {
	var validaterErr = validators.ValidateServer(server)
	if validaterErr != nil {
		return "", validaterErr
	}

	var command = ""

	if server.NoSudo {
		command = "apt update"
	} else {
		command = "echo " + server.Password + "| sudo -S apt update"
	}

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

	var command = "";

	if server.NoSudo {
		command = "apt list --upgradable 2>/dev/null | grep upgradable | wc -l"
	} else {
		command = "echo " + server.Password + "| sudo -S apt list --upgradable 2>/dev/null | grep upgradable | wc -l"
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

func GetUpgrades(server config.Server) (bool, error) {
	var _, updateCacheErr = UpdateCache(server)
	if updateCacheErr != nil {
		return false, updateCacheErr;
	}
	var hasUpdate, _, upgradableErr = GetUpgradable(server)
	if upgradableErr != nil {
		return false, upgradableErr;
	}
	if hasUpdate {
		return true, nil;
	}
	return false, nil;
}