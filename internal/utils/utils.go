package utils

import (
	"strings"

	"github.com/steveiliop56/puck/internal/config"
	"github.com/steveiliop56/puck/internal/ssh"
)

func GetDistro(server config.Server) (string, error) {
	var sshOutput, sshErr = ssh.RunCommandRich(server, `cat /etc/os-release | grep "^ID=" | awk '{print $1}' | cut -d "=" -f2 | cut -d '"' -f2`)
	if sshErr != nil {
		return "", sshErr
	}
	return sshOutput, nil
}

func GetCommand(distro string) ([]string, bool) {
	var commands []string
	var skip = false;
	// Need to test this
	switch strings.TrimSpace(distro) {
	case "ubuntu": 
		commands = append(commands, "apt update")
		commands = append(commands, "apt list --upgradable 2>/dev/null | grep upgradable | wc -l")
	case "debian":
		commands = append(commands, "apt update")
		commands = append(commands, "apt list --upgradable 2>/dev/null | grep upgradable | wc -l")
	case "fedora":
		commands = append(commands, "dnf check-update")
		commands = append(commands, `dnf check-update | grep -c '^\S'`)
	case "opensuse-leap":
		commands = append(commands, "dnf check-update")
		commands = append(commands, `dnf check-update | grep -c '^\S'`)
	case "alpine":
		commands = append(commands, "apk update")
		commands = append(commands, "apk version -u | wc -l")
	case "arch":
		commands = append(commands, "pacman -Syyu")
		commands = append(commands, "pacman -Qu | wc -l")
	default:
		skip = true
	}
	return commands, skip
}