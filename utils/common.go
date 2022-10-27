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

func ProcessIsExists(pid string) (bool, error) {
	cmdStr := fmt.Sprintf("kill -s 0 %s", pid)
	bytes, err := ExecWithOut(cmdStr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if len(bytes) == 0 {
		return true, err
	}
	return false, err
}
