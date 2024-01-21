package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func runCommandShell(command string) error {
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

func RunCommand(commands []string) error {
	return runCommandShell(strings.Join(commands, " && "))
}

func RunOnlineScript(urlStr string) error {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	fileName := "." + path.Base(parsedUrl.Path)

	output, err := os.Create(fileName)
	if err != nil {
		return err
	}

	output.Chmod(0755)
	defer output.Close()

	response, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	err = RunCommand([]string{"./" + fileName})
	os.Remove(fileName)

	return err
}
