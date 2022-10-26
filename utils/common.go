package utils

import (
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
