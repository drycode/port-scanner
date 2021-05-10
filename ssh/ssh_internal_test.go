package ssh

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	. "github.com/drypycode/portscanner/utils"
	env "github.com/joho/godotenv"
)

var err error = env.Load("../.env")

var TEST_CONF SSHConfig = SSHConfig{
	Key: os.Getenv("SSH_KEY_PATH"), User: os.Getenv("REMOTE_USER"), RemoteHost: os.Getenv("REMOTE_HOST"), Port: "22",
}

func TestInspectOS(t *testing.T) {
	check(err)
	session := setupSession(TEST_CONF)
	result := inspectRemoteOS("hostnamectl", session)
	result = bytes.Trim(result, "\x00")
	result = append(result, []byte("\n")...)
	arch := getValue(result, "Architecture")
	os := getValue(result, "Operating System")

	AssertEquals(t, "", "arm64", arch)
	AssertEquals(t, "", "Amazon Linux 2", os)
}

func TestRunCommand(t *testing.T) {
	check(err)
	session := setupSession(TEST_CONF)
	cmd := "./main --ports=80,4423,100-105,40-45,1000-1200 --hosts='127.0.0.1,localhost,google.com' --output=/tmp/dat2.json --protocol=TCP"
	runCommand(cmd, session)
}

// func TestBuildPSOnRemote(t *testing.T) {
// 	buildPSOnRemote(TEST_CONF)
// }

// CURRENTLY FAILING BECAUSE os.Getcwd is very fragile
// func TestJump(t *testing.T) {
// 	os.Args = []string{"go", "run", "main.go", "--ports=80,4423,100-105,40-45,1000-2000", "--hosts='127.0.0.1,localhost,google.com'", "--output=/tmp/dat2.json", "--protocol=TCP"}
// 	Jump(TEST_CONF)
// }

func TestRemoveRemoteFlags(t *testing.T) {
	args := []string{
		"go", "run", "main.go", "--ports=80,4423,100-105,40-45,1000-2000",
		"--hosts='127.0.0.1,localhost,google.com'", "--output=/tmp/dat2.json",
		"--protocol=TCP", fmt.Sprintf("--remote-host=%s", os.Getenv("REMOTE_HOST")),
		"--remote-user=ec2-user", fmt.Sprintf("--ssh-key=%s", os.Getenv("SSH_KEY_PATH")),
		"--jump",
	}
	removeRemoteFlags(args)
	fmt.Println(args)
	AssertEquals(t, "Removed elements", []string{
		"go", "run", "main.go", "--ports=80,4423,100-105,40-45,1000-2000",
		"--hosts='127.0.0.1,localhost,google.com'", "--output=/tmp/dat2.json", "--protocol=TCP", "", "", "", "",
	}, args)
}
