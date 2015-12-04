package SSHClient

import (
	"fmt"
	"bytes"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Host   string
	Port   int
	Config *ssh.ClientConfig
	Cmd *SSHCommand
}

type SSHCommand struct {
	Cmd    string
	Stdout string
	Stderr string
	Err    error
}

// Create a new ssh.Session in the SSHClient
func (client *SSHClient) newSession() (*ssh.Session, error) {
	conn_string := fmt.Sprintf("%s:%d", client.Host, client.Port)
	connection, err := ssh.Dial("tcp", conn_string, client.Config)
	if err != nil {

		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {

		return nil, fmt.Errorf("Failed to create session: %s", err)
	}

	modes := ssh.TerminalModes{
		// ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		return nil, fmt.Errorf("Request for pseudo terminal failed: %s", err)
	}

	return session, nil
}

// Create a new ssh.Session from the SSHClient and execute the SSHCommand
func ExecuteCmd(client *SSHClient) bool {
	session, _ := client.newSession()
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	var stderrBuf bytes.Buffer
	session.Stderr = &stderrBuf

	err := session.Run(client.Cmd.Cmd)

	client.Cmd.Stdout = stdoutBuf.String()
	client.Cmd.Stderr = stderrBuf.String()
	client.Cmd.Err = err

		if err == nil {
			return true
		} else {
			return false
		}
}

// Return a ssh.AuthMethod from a keypair filepath
func PublicKeyFile(filepath string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//
// func ExecuteCmd(client *SSHClient, cmd *SSHCommand) bool {
// 	session, _ := client.NewSession()
// 	defer session.Close()
//
// 	var stdoutBuf bytes.Buffer
// 	session.Stdout = &stdoutBuf
// 	var stderrBuf bytes.Buffer
// 	session.Stderr = &stderrBuf
//
// 	err := session.Run(cmd.Cmd)
//
// 	cmd.Stdout = stdoutBuf.String()
// 	cmd.Stderr = stderrBuf.String()
// 	cmd.Err = err
//
// 	if err == nil {
// 		return true
// 	} else {
// 		return false
// 	}
// }
