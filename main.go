package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printQrepoBanner()
		printCommandHelp()
		return
	}

	switch os.Args[1] {
	case "init":
		printQrepoBanner()
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

	case "log":
		printQrepoBanner()
		if len(os.Args) != 2 {
			printLogHelp()
			return
		}

		logRepo()

	case "x":
		executeOnlineScript()

	default:
		printQrepoBanner()
		printCommandHelp()
	}
}
