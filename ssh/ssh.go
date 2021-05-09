package ssh

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var GOOS map[string]string = map[string]string{"Amazon Linux 2": "linux"}
var GOARCH map[string]string = map[string]string{"arm64": "arm64"}

var REMOTE_GOOS string
var REMOTE_GOARCH string
var SHARED_SESSION *ssh.Session
var BINARY_NAME string = "main"

type SSHConfig struct {
	Key        string
	User       string
	RemoteHost string
	Port       string
}

func removeRemoteFlags(args []string) {
	for i, arg := range args {
		for _, str := range []string{"--remote-host", "--jump", "--remote-user", "--ssh-key"} {
			if strings.Contains(arg, str) {
				args[i] = ""
				break
			}
		}
	}

}

func Jump(conf SSHConfig) {
	session1 := setupSession(conf)
	defer session1.Close()
	session2 := setupSession(conf)
	removeRemoteFlags(os.Args)
	os.Args[0] = fmt.Sprintf("./%s", BINARY_NAME)
	cmd := strings.Join(os.Args, " ")
	logrus.Info(cmd)
	// TODO:  - this will grep the terminal command, remove the -i and --jump flags from teh command, and
	if strings.Compare(os.Args[0], "go") == 0 {
		cmd = fmt.Sprintf("./%s", BINARY_NAME) + " " + strings.Join(os.Args[3:], " ")
	}

	// TODO: - if the -o flag is given, we should scp the output file back to the local machine
	buildPSOnRemote(conf, session1)
	runCommand(cmd, session2)

}

func setupSession(conf SSHConfig) *ssh.Session {
	config := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			publicKey(conf.Key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", conf.RemoteHost+":"+conf.Port, config)
	check(err)

	sess, err := conn.NewSession()
	check(err)

	return sess
}

func getValue(output []byte, key string) string {
	keyStart := bytes.Index(output, []byte(key))
	output = output[keyStart:]
	delimiter := bytes.Index(output, []byte(":"))
	output = output[delimiter:]
	end := bytes.Index(output[2:], []byte("\n"))
	return string(output[2 : end+2])
}

func pipeTerminalOut(sess *ssh.Session) {
	sessStdout, err := sess.StdoutPipe()
	check(err)
	sessStderr, err := sess.StderrPipe()
	check(err)

	// Incase the remote container has some issue
	go io.Copy(os.Stderr, sessStderr)
	go io.Copy(os.Stdout, sessStdout)
}

func inspectRemoteOS(cmd string, sess *ssh.Session) []byte {
	sessStdout, err := sess.StdoutPipe()
	check(err)
	sessStderr, err := sess.StderrPipe()
	check(err)

	// Incase the remote container has some issue
	go io.Copy(os.Stderr, sessStderr)

	err = sess.Run(cmd)

	check(err)
	b := make([]byte, 2000)
	sessStdout.Read(b)
	REMOTE_GOOS = getValue(b, "Operating System")
	REMOTE_GOARCH = getValue(b, "Architecture")
	return b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	check(err)
	signer, err := ssh.ParsePrivateKey(key)
	check(err)
	return ssh.PublicKeys(signer)
}

func runCommand(cmd string, sess *ssh.Session) {
	pipeTerminalOut(sess)
	fmt.Println("114: " + cmd)
	err := sess.Run(cmd)

	check(err)
}

// func checkRemoteFileExists(remotePath string) {

// }

func buildPSOnRemote(conf SSHConfig, sess *ssh.Session) {
	var remoteBuildPath string = fmt.Sprintf(":/home/%s/main", conf.User)
	// TODO: checkRemoteFileExists(remoteBuildPath)

	scpPath := func(tmpDir string) string {
		command := fmt.Sprintf("scp -i %s ", conf.Key)
		command = command + filepath.Join(tmpDir, "main") + " " + fmt.Sprintf("%s@%s", conf.User, conf.RemoteHost)
		command = command + remoteBuildPath
		return command
	}
	// Check if build is there already
	dir, err := ioutil.TempDir("/tmp", "portscanner_build")
	// fmt.Println(dir)
	check(err)

	fileName, err := os.Getwd()
	check(err)
	// fileName = filepath.Join(fileName, "../")
	// fmt.Println(fileName)
	inspectRemoteOS("hostnamectl", sess)
	mainPath := filepath.Join(fileName, "main.go")
	fmt.Println(mainPath)

	_, err = os.Stat(dir)
	check(err)
	_, err = os.Stat(mainPath)
	check(err)

	buildCommand := fmt.Sprintf(
		"env GOOS=%s GOARCH=%s go build -o %s %s",
		GOOS[REMOTE_GOOS], GOARCH[REMOTE_GOARCH], dir,
		mainPath,
	)
	scpCommand := scpPath(dir)

	// fmt.Println(buildCommand)
	// fmt.Println(scpCommand)
	bc := strings.Split(buildCommand, " ")
	comm := exec.Command(bc[0], bc[1:]...)
	err = comm.Run()
	check(err)

	scpc := strings.Split(scpCommand, " ")
	comm = exec.Command(scpc[0], scpc[1:]...)
	err = comm.Run()
	check(err)

	defer os.RemoveAll(dir)
}
