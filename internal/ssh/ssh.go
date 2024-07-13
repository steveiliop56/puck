package ssh

import (
	"bytes"
	"net"
	"os"

	"github.com/steveiliop56/puck/internal/config"
	"golang.org/x/crypto/ssh"
)

// Given the ssh credentials it runs the command
func RunCommand(hostname string, username string, password string, privateKey string, command string) (string, error) {
	var authConfig []ssh.AuthMethod

	if privateKey == "" {
		authConfig = []ssh.AuthMethod{
            ssh.Password(password),
        }
	} else {
		key, err := ssh.ParsePrivateKey([]byte(privateKey))
		if err != nil {
			return "", err
		}
		authConfig = []ssh.AuthMethod{
            ssh.PublicKeys(key),
        }
	}

	sshConfig := &ssh.ClientConfig{
		User: username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: authConfig,
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(hostname, "22"), sshConfig)
	if err != nil {
		return "", err
	}

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	var b bytes.Buffer

	session.Stdout = &b

	err = session.Run(command)

	return b.String(), err
}

// Given the server struct and the command it just check for the private key and uses it too but other than that it just returns the ssh output
func RunCommandRich(server config.Server, command string) (string, error) {
	var privateKey = ""

	if server.PrivateKey != "" {
		var content, readErr = os.ReadFile(server.PrivateKey)
		if readErr != nil {
			return "", readErr
		}
		privateKey = string(content)
	}

	var sshOutput, sshErr = RunCommand(server.Hostname, server.Username, server.Password, privateKey, command)

	if sshErr != nil {
		return "", sshErr
	}

	return sshOutput, nil
}
