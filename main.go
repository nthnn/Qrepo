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
package main

import (
	"os"

	"github.com/nthnn/Qrepo/banner"
	"github.com/nthnn/Qrepo/core"
)

func main() {
	if len(os.Args) == 1 {
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

	case "update":
		core.UpdateQrepo()

	default:
		banner.PrintQrepoBanner()
		banner.PrintCommandHelp()
	}
}
