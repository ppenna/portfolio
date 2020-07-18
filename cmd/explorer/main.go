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
	"fmt"
	"os"
	"os/exec"
	"portfolio/internal/asset"
	"portfolio/internal/config"
	"portfolio/internal/database"
	"strconv"
)

var plotConfigs = map[string]int{
	database.TickerALZR11: 5,
	database.TickerBPFF11: 10,
	database.TickerBRCR11: 10,
	database.TickerHGBS11: 5,
	database.TickerHGFF11: 15,
	database.TickerHGLG11: 5,
	database.TickerHGRE11: 5,
	database.TickerJSRE11: 10,
	database.TickerKNCR11: 25,
	database.TickerKNIP11: 25,
	database.TickerVISC11: 10,
	database.TickerXPLG11: 10,
	database.TickerXPML11: 3,
}

// Plot charts on an asset.
func plotChart(a *asset.Asset) {
	ylim, ok := plotConfigs[a.Ticker()]
	if !ok {
		fmt.Println("missing plot configuration")
	}

	cmd := exec.Command("Rscript",
		"--vanilla",
		config.RplotsPath+"pb.R",
		a.Ticker(),
		strconv.Itoa(ylim),
	)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s", stdoutStderr)
}

// Print information on all assets.
func printAllAssets() {
	for _, a := range database.Assets() {
		a.Write(os.Stdout)

		// Plot charts.
		if plotCharts {
			plotChart(a)
		}
	}
}

// Print information on a target asset.
func printAsset(ticker string) {
	a, err := database.GetAssetByTicker(ticker)
	if err != nil {
		fmt.Println(err)
	}

	a.Write(os.Stdout)

	// Plot charts.
	if plotCharts {
		plotChart(a)
	}
}

// List all assets.
func listAssets() {
	fmt.Println("Known Assets:")
	for _, a := range database.Assets() {
		fmt.Println("  " + a.Ticker())
	}

}

func main() {

	parseArgs()

	database.Load()

	// Parse command.
	if printAll {
		printAllAssets()
	} else if list {
		listAssets()
	} else {
		printAsset(ticker)
	}
}
