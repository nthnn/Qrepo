package main

import (
	"os"

	"github.com/nthnn/Qrepo/banner"
	"github.com/nthnn/Qrepo/core"
)

func main() {
	if len(os.Args) < 2 {
		banner.PrintQrepoBanner()
		banner.PrintCommandHelp()
		return
	}

	switch os.Args[1] {
	case "init":
		banner.PrintQrepoBanner()
		if len(os.Args) != 2 {
			banner.PrintInitHelp()
			return
		}

		core.InitRepo()

	case "run":
		if len(os.Args) != 3 {
			banner.PrintRunHelp()
			return
		}

		core.RunScript(os.Args[2])

	case "log":
		banner.PrintQrepoBanner()
		if len(os.Args) != 2 {
			banner.PrintLogHelp()
			return
		}

		core.LogRepo()

	case "x":
		core.ExecuteOnlineScript()

	default:
		banner.PrintQrepoBanner()
		banner.PrintCommandHelp()
	}
}
