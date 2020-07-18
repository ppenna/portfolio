/*
 * MIT License
 *
 * Copyright(c) 2020 Pedro Henrique Penna <pedrohenriquepenna@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"flag"
)

// Command Line Arguments
var (
	saveWallet     bool   // Save wallet?
	walletFilename string // Wallet File name
	printStats     bool   // Print statistics?
	printWallet    bool   // Print wallet?
)

// Parses command line arguments.
func parseArgs() {

	saveHelp := "Save wallet to a file?"
	flag.BoolVar(&saveWallet, "save", false, saveHelp)

	walletFilenameHelp := "Name of the wallet file"
	flag.StringVar(&walletFilename, "output", "new.wallet", walletFilenameHelp)

	printStatsHelp := "Print statistics?"
	flag.BoolVar(&printStats, "stats", true, printStatsHelp)

	printWalletHelp := "Print wallet?"
	flag.BoolVar(&printWallet, "print", false, printWalletHelp)

	flag.Parse()
}
