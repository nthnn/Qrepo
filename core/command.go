/*
 * This file is part of the Qrepo.
 * Copyright (c) 2024 Nathanne Isip
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package core

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/nthnn/Qrepo/util"
	"golang.org/x/mod/semver"
)

func InitRepo() {
	if hasQrepoJsonAlready() {
		fmt.Println("Already \033[31minitialized\033[0m Qrepo project.")
		return
	}

	projectName, err := os.Getwd()
	if err != nil {
		projectName = "prj_" + util.GetCurrentUnixTimeAsString()
	} else {
		paths := strings.Split(projectName, string(os.PathSeparator))
		projectName = paths[len(paths)-1]
	}

	fmt.Print("Project name (" + projectName + "): ")
	tmpName := util.ReadString()
	if tmpName != "" {
		projectName = tmpName
	}

	projectAuthor, err := util.GetCurrentUsername()
	if err != nil {
		projectAuthor = "anonymous"
	}

	fmt.Print("Author (" + projectAuthor + "): ")
	tmpAuthor := util.ReadString()

	if tmpAuthor != "" {
		projectAuthor = tmpAuthor
	}

	fmt.Print("Version (v0.0.1): ")
	projectVersion := util.ReadString()

	if projectVersion == "" {
		projectVersion = "v0.0.1"
	}

	if !semver.IsValid(projectVersion) {
		fmt.Println("\033[31mError\033[0m: Invalid sematic versioning string.")
		return
	}

	fmt.Print("Description: ")
	projectDesc := util.ReadString()

	gitOrigin, err := util.GetGitRemoteOrigin()
	if err != nil || gitOrigin == "" {
		gitOrigin = "null"
	}

	fmt.Print("Git (" + gitOrigin + "): ")
	tmpOrigin := util.ReadString()

	if tmpOrigin != "" {
		gitOrigin = tmpOrigin
	}

	if _, err := url.ParseRequestURI(gitOrigin); err != nil {
		fmt.Println("\033[31mError\033[0m: Invalid Git origin URL.")
		return
	}

	generated := "{\n\t\"name\": \"" + projectName + "\"," +
		"\n\t\"author\": \"" + projectAuthor + "\"," +
		"\n\t\"version\": \"" + projectVersion + "\"," +
		"\n\t\"description\": \"" + projectDesc + "\"," +
		"\n\t\"git\": \"" + gitOrigin + "\"," +
		"\n\t\"scripts\": {\n\t\t\"test\": [\"echo No test specified.\"]\n\t}" +
		"\n}"

	if err := util.WriteStringToFile("qrepo.json", generated); err != nil {
		fmt.Println("\033[31mError\033[0m: " + err.Error())
		return
	}

	fmt.Println("\nQrepo \033[32msuccessfully\033[0m initialized!")
}

func RunScript(scriptName string) {
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

	util.RunCommand(qrepo.Scripts[key])
}

func ExecuteOnlineScript() {
	if len(os.Args) != 3 {
		fmt.Println("A script URL \033[31mwas not\033[0m specified.")
		return
	}

	if err := util.RunOnlineScript(os.Args[2]); err != nil {
		fmt.Println("\033[31mError: \033[0m" + err.Error())
	}
}

func UpdateQrepo() {
	if err := util.RunOnlineScript("https://raw.githubusercontent.com/nthnn/Qrepo/master/support/install.sh"); err != nil {
		fmt.Println("\033[31mError: \033[0m" + err.Error())
		os.Exit(0)
	}

	fmt.Println("[\033[92m+\033[0m] Done updating \033[36mQrepo\033[0m!")
}

func LogRepo() {
	if !hasQrepoJsonAlready() {
		fmt.Println("\033[31mNot\033[0m a Qrepo repository project.")
		return
	}

	qrepo, err := extractQrepoInfos()
	if err != nil {
		fmt.Println("\033[31mError\033[0m: " + err.Error())
		return
	}

	fmt.Println("\033[35mName\033[0m:\t\t" + qrepo.Name)
	fmt.Println("\033[35mAuthor\033[0m:\t\t" + qrepo.Author)
	fmt.Println("\033[35mVersion\033[0m:\t" + qrepo.Version)
	fmt.Println("\033[35mDescription\033[0m:\t" + qrepo.Description)
	fmt.Println("\033[35mGit\033[0m:\t\t" + qrepo.Git)
	fmt.Println("\033[35mScripts\033[0m:\t" + strings.Join(qrepo.getMapKeys(), ", "))
}
