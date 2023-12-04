package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func readString() string {
	ln := ""
	fmt.Scanf("%s", &ln)

	return ln
}

func runCommand(command string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	default:
		cmd = exec.Command("sh", "-c", command)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	return nil
}
