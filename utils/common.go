package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func ExecWithOut(cmdStr string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	closer, err := cmd.StdoutPipe()
	defer func() {
		_ = closer.Close()
		_ = cmd.Wait()
	}()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(closer)
}

func GetProcessID(procArgs []string) (string, error) {
	cmdStr := "ps -ef"
	for _, arg := range procArgs {
		cmdStr += fmt.Sprintf(" | grep %s", arg)
	}
	cmdStr += " | grep -v grep | awk '{print $2}'"
	bytes, err := ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
