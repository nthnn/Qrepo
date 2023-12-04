package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func readString() string {
	ln := ""
	fmt.Scanf("%s", &ln)

	return ln
}

func writeStringToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func hasQrepoJsonAlready() bool {
	_, err := os.Stat("qrepo.json")
	return !os.IsNotExist(err)
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

func getCurrentUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}

func getCurrentUnixTimeAsString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func getGitRemoteOrigin() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	remoteURL := strings.TrimSpace(string(output))
	return remoteURL, nil
}
