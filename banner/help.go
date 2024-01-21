package banner

import "fmt"

func PrintCommandHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo [command] <arguments> ...")

	fmt.Println("\nCommands:")
	fmt.Println("\tinit\tInitialize current dir as repository project.")
	fmt.Println("\trun\tRun defined script in the project folder.")
	fmt.Println("\tlog\tLogs the informations on the current repository.")
	fmt.Println("\tx\tExecute a script from specified URL.")
	fmt.Println("\tupdate\tUpdates the current installed Qrepo.")
	fmt.Println("\thelp\tShow help message for specific command.")

	fmt.Println()
}

func PrintInitHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo init")

	fmt.Println("\nDescription:")
	fmt.Println("\tInitializes the current working directory")
	fmt.Println("\tas new Qrepo repository project and generates")
	fmt.Println("\t`qrepo.json` file in the current folder.")

	fmt.Println()
}

func PrintRunHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo run <script name>")

	fmt.Println("\nParameter:")
	fmt.Println("\tscript name - Name of script to be executed.")

	fmt.Println("\nDescription:")
	fmt.Println("\tRuns a script defined in the `qrepo.json` file.")
	fmt.Println()
}

func PrintLogHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo log")

	fmt.Println("\nDescription:")
	fmt.Println("\tLogs the informations of the current Qrepo;")
	fmt.Println("\tincluding the name of the project, author,")
	fmt.Println("\tgit origin, and the available scripts.")
	fmt.Println()
}
