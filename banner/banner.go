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

const qrepoBanner = "\n \033[32m██████\033[0m\033[33m╗\033[0m \033[32m██████" +
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

func PrintQrepoBanner() {
	fmt.Println(qrepoBanner)
	fmt.Println("\n\033[97m\033[44m       Qrepo Package Manager v1.5.3       \033[0m\033[0m")
	fmt.Println()
}
