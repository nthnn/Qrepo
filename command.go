package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func initRepo() {
	if hasQrepoJsonAlready() {
		fmt.Println("Already \033[31minitialized\033[0m Qrepo project.")
		return
	}

	projectName, err := os.Getwd()
	if err != nil {
		projectName = "prj_" + getCurrentUnixTimeAsString()
	} else {
		paths := strings.Split(projectName, string(os.PathSeparator))
		projectName = paths[len(paths)-1]
	}

	fmt.Print("Project name (" + projectName + "): ")
	tmpName := readString()
	if tmpName != "" {
		projectName = tmpName
	}

	projectAuthor, err := getCurrentUsername()
	if err != nil {
		projectAuthor = "anonymous"
	}

	fmt.Print("Author (" + projectAuthor + "): ")
	tmpAuthor := readString()

	if tmpAuthor != "" {
		projectAuthor = tmpAuthor
	}

	gitOrigin, err := getGitRemoteOrigin()
	if err != nil || gitOrigin == "" {
		gitOrigin = "null"
	}

	fmt.Print("Git (" + gitOrigin + "): ")
	tmpOrigin := readString()

	if tmpOrigin != "" {
		gitOrigin = tmpOrigin
	}

	generated := "{\n\t\"name\": \"" + projectName + "\"," +
		"\n\t\"author\": \"" + projectAuthor + "\"," +
		"\n\t\"git\": \"" + gitOrigin + "\"," +
		"\n\t\"scripts\": {\n\t\t\"test\": [\"echo No test specified.\"]\n\t}" +
		"\n}"

	if err := writeStringToFile("qrepo.json", generated); err != nil {
		fmt.Println("Error occured: " + err.Error())
		return
	}

	fmt.Println("\nQrepo \033[32msuccessfully\033[0m initialized!")
}

func runScript(scriptName string) {
	if !hasQrepoJsonAlready() {
		fmt.Println("\033[31mNot\033[0m a Qrepo repository project.")
		return
	}

	qrepo, err := extractQrepoInfos()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	os := runtime.GOOS
	arch := runtime.GOARCH

	key := os + "/" + arch + ":" + scriptName
	if _, ok1 := qrepo.Scripts[key]; !ok1 {
		key = os + ":" + scriptName

		if _, ok2 := qrepo.Scripts[key]; !ok2 {
			key = scriptName
		}
	}

	if _, ok := qrepo.Scripts[key]; !ok {
		fmt.Println("Script \"" + scriptName + "\" \033[31mnot\033[0m found.")
		return
	}

	runCommand(qrepo.Scripts[key])
}

func logRepo() {
	if !hasQrepoJsonAlready() {
		fmt.Println("\033[31mNot\033[0m a Qrepo repository project.")
		return
	}

	qrepo, err := extractQrepoInfos()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Println("\033[35mName\033[0m:\t\t" + qrepo.Name)
	fmt.Println("\033[35mAuthor\033[0m:\t\t" + qrepo.Author)
	fmt.Println("\033[35mGit\033[0m:\t\t" + qrepo.Git)
	fmt.Println("\033[35mScripts\033[0m:\t" + strings.Join(getMapKeys(qrepo.Scripts), ", "))
}
