package ssh

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

type SSHConfig struct {
	key        string
	user       string
	remoteHost string
	command    string
}

func Tunnel(conf SSHConfig) {

	config := &ssh.ClientConfig{
		User: conf.user,
		Auth: []ssh.AuthMethod{
			publicKey(conf.key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", conf.remoteHost, config)
	check(err)
	defer conn.Close()
	runCommand(conf.command, conn)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func runCommand(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	check(err)
	defer sess.Close()

	sessStdOut, err := sess.StdoutPipe()
	check(err)

	sessStderr, err := sess.StderrPipe()
	check(err)

	go io.Copy(os.Stdout, sessStdOut)
	check(err)

	go io.Copy(os.Stderr, sessStderr)

	err = sess.Run(cmd)

	check(err)
}
