package util

import (
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func GetCurrentUnixTimeAsString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GetCurrentUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}

func GetGitRemoteOrigin() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	remoteURL := strings.TrimSpace(string(output))
	return remoteURL, nil
}
