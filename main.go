package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	printQrepoBanner()

	if len(os.Args) < 2 {
		printCommandHelp()
		return
	}

	switch os.Args[1] {
	case "init":
		if len(os.Args) != 2 {
			printInitHelp()
			return
		}

		initRepo()

	case "run":
		if len(os.Args) != 3 {
			printRunHelp()
			return
		}

		runScript(os.Args[2])

	case "help":
		printCommandHelp()
	}
}

func initRepo() {
	if hasQrepoJsonAlready() {
		fmt.Println("Already initialized a Qrepo project.")
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
		"\n\t\"scripts\": {\n\t\t\"test\": \"echo No test specified.\"\n\t}" +
		"\n}"

	if err := writeStringToFile("qrepo.json", generated); err != nil {
		fmt.Println("Error occured: " + err.Error())
		return
	}

	fmt.Println("\nQrepo \033[32msuccessfully\033[0m initialized!")
}

func runScript(scriptName string) {

}
