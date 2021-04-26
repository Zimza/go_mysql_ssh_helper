package internal

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

func getKeyFile(sshKeyPath string) (key ssh.Signer, err error) {
	buf, err := ioutil.ReadFile(sshKeyPath)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}

// SSHClient : Return SSH Client
func SSHClient(sshHost *string, sshUser *string, sshKeyPath *string, sshPort *string) (sshtun *ssh.Client) {
	// Now in the main function DO:
	key, err := getKeyFile(*sshKeyPath)
	if err != nil {
		panic(err)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User: *sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 10,
	}

	// Connect to the SSH Server
	sshtun, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", *sshHost, *sshPort), sshConfig)
	if err != nil {
		panic(err)
	}

	return sshtun
}
