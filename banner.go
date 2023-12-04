package main

import "fmt"

const QrepoBanner = "\n \033[32m██████\033[0m\033[33m╗\033[0m \033[32m██████" +
	"\033[0m\033[33m╗\033[0m \033[32m███████\033[0m\033[33m╗\033[0m\033[32m" +
	"██████\033[0m\033[33m╗\033[0m  \033[32m██████\033[0m\033[33m╗\033[0m\n" +
	"\033[32m██\033[0m\033[33m╔═══\033[0m\033[32m██\033[0m\033[33m╗\033[0m" +
	"\033[32m██\033[0m\033[33m╔══\033[0m\033[32m██\033[0m\033[33m╗\033[0m" +
	"\033[32m██\033[0m\033[33m╔════╝\033[0m\033[32m██\033[0m\033[33m╔══\033[0m" +
	"\033[32m██\033[0m\033[33m╗\033[0m\033[32m██\033[0m\033[33m╔═══\033[0m" +
	"\033[32m██\033[0m\033[33m╗\033[0m\n\033[32m██\033[0m\033[33m║\033[0m   " +
	"\033[32m██\033[0m\033[33m║\033[0m\033[32m██████\033[0m\033[33m╔╝\033[0m" +
	"\033[32m█████\033[0m\033[33m╗\033[0m  \033[32m██████\033[0m\033[33m╔╝" +
	"\033[0m\033[32m██\033[0m\033[33m║\033[0m   \033[32m██\033[0m\033[33m║" +
	"\033[0m\n\033[32m██\033[0m\033[33m║\033[0m\033[32m▄▄ ██\033[0m\033[33m║" +
	"\033[0m\033[32m██\033[0m\033[33m╔══\033[0m\033[32m██\033[0m\033[33m╗" +
	"\033[0m\033[32m██\033[0m\033[33m╔══╝\033[0m  \033[32m██\033[0m\033[33m" +
	"╔═══╝\033[0m \033[32m██\033[0m\033[33m║\033[0m   \033[32m██\033[0m" +
	"\033[33m║\033[0m\n\033[33m╚\033[0m\033[32m██████\033[0m\033[33m╔╝\033[0m" +
	"\033[32m██\033[0m\033[33m║\033[0m  \033[32m██\033[0m\033[33m║\033[0m" +
	"\033[32m███████\033[0m\033[33m╗\033[0m\033[32m██\033[0m\033[33m║\033[0m" +
	"     \033[33m╚\033[0m\033[32m██████\033[0m\033[33m╔╝\033[0m\n \033[33m╚══" +
	"\033[0m\033[32m▀▀\033[0m\033[33m═╝ ╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝\033[0m"

func printQrepoBanner() {
	fmt.Println(QrepoBanner)
	fmt.Println("\n\033[97m\033[44m       Qrepo Package Manager v1.0.0       \033[0m\033[0m")
	fmt.Println()
}

func printCommandHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo [command] <arguments> ...")

	fmt.Println("\nCommands:")
	fmt.Println("\tinit\tInitialize current dir as repository project.")
	fmt.Println("\trun\tRun defined script in the project folder.")
	fmt.Println("\tlog\tLogs the informations on the current repository.")
	fmt.Println("\thelp\tShow help message for specific command.")

	fmt.Println()
}

func printInitHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo init")

	fmt.Println("\nDescription:")
	fmt.Println("\tInitializes the current working directory")
	fmt.Println("\tas new Qrepo repository project and generates")
	fmt.Println("\t`qrepo.json` file in the current folder.")

	fmt.Println()
}

func printRunHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo run <script name>")

	fmt.Println("\nParameter:")
	fmt.Println("\tscript name - Name of script to be executed.")

	fmt.Println("\nDescription:")
	fmt.Println("\tRuns a script defined in the `qrepo.json` file.")
	fmt.Println()
}

func printLogHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tqrepo log")

	fmt.Println("\nDescription:")
	fmt.Println("\tLogs the informations of the current Qrepo;")
	fmt.Println("\tincluding the name of the project, author,")
	fmt.Println("\tgit origin, and the available scripts.")
}
