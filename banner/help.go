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
