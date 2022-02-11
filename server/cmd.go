package main

import (
	"fmt"
	"os/exec"
)

func RunCmd(cmdStr string, args ...string) error {
	cmd := exec.Command(cmdStr, args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return err
	} else {
		fmt.Println(string(out))
	}
	return nil
}
