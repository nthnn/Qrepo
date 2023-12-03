package main

import (
	"os"
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

	case "run":
		if len(os.Args) != 3 {
			printRunHelp()
			return
		}

	case "help":
		printCommandHelp()
	}
}
