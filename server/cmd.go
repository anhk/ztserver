package main

import (
	"os/exec"
)

func RunCmd(cmdStr string, args ...string) error {
	return exec.Command(cmdStr, args...).Run()
}
