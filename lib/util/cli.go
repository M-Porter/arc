package util

import (
	"fmt"
	"os/exec"
)

func EnforceFlag(flag interface{}, compare interface{}, message string) {
	if flag == compare {
		Fatalf(message)
	}
}

func RunCmd(command string) ([]byte, error) {
	return exec.Command("bash", "-c", command).Output()
}

func RunCmdAt(command string, path string) ([]byte, error) {
	return RunCmd(fmt.Sprintf("cd %s && %s", path, command))
}
