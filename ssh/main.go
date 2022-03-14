package main

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var username = "root"
var privateKeyFile = "/Users/gerardyap/.ssh/id_rsa"
var host = "35.189.188.149:22"
var commandToExecute = "hostname"

func main() {
	privateKey := getKeySigner(privateKeyFile)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	// Multiple sessions per client are allowed
	// A session can only perform one action
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Pipe the session output directly to standard output
	// Thanks to the convenience of writer interface
	session.Stdout = os.Stdout

	err = session.Run(commandToExecute)
	if err != nil {
		log.Fatal("Error executing command. ", err)
	}

	// Start a new session
	session, err = client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	// Pipe the standard buffers together
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	// Get psuedo-terminal
	err = session.RequestPty(
		"linux", // or "vt100", "xterm"
		40,      // Height
		80,      // Width
		// https://godoc.org/golang.org/x/crypto/ssh#TerminalModes
		// POSIX Terminal mode flags defined in RFC 4254 Section 8.
		// https://tools.ietf.org/html/rfc4254#section-8
		ssh.TerminalModes{
			ssh.ECHO: 0,
		})
	if err != nil {
		log.Fatal("Error requesting psuedo-terminal. ", err)
	}

	// Run shell until it is exited
	err = session.Shell()
	if err != nil {
		log.Fatal("Error executing command. ", err)
	}
	session.Wait()
}

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("Error loading private key file. ", err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatal("Error parsing private key. ", err)
	}
	return privateKey
}
