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

package asset

import (
	"fmt"
	"os"
)

// Asset
type Asset struct {
	id     int              // ID
	ticker string           // Ticker
	hist   *AssetHistory    // Historical Data
	stats  *AssetStatistics // Statistics
	class  int              // Class
}

/*============================================================================*
 * read()                                                                     *
 *============================================================================*/

// Reads an asset from a file.
func Read(id int, ticker, filename string, class int) *Asset {
	a := &Asset{}

	// TODO: assert arguments.

	a.id = id
	a.ticker = ticker
	a.class = class
	a.hist = readHistory(filename)
	a.stats = computeStatistics(a.hist)

	return a
}

/*============================================================================*
 * Utilities                                                                  *
 *============================================================================*/

// Normalizes x according to y.
func normalize(x, y float32) float32 {
	return (x - y) / y
}

/*============================================================================*
 * Getters                                                                    *
 *============================================================================*/

// Returns the ID of the target asset.
func (a *Asset) ID() int { return a.id }

// Returns the ticker of the target asset.
func (a *Asset) Ticker() string { return a.ticker }

// Returns the class of the target asset.
func (a *Asset) Class() int { return a.class }

// Returns the performance of the target asset.
func (a *Asset) Performance() float32 {
	return a.stats.aagrSharePrice + a.stats.emaDY
}

// Returns the risk of the target asset.
func (a *Asset) Risk() float32 {
	return 0
}

// Returns the valuation of the target asset.
func (a *Asset) Cost() float32 {
	var cost float32
	var div float32

	div = 2.0
	cost += -normalize(a.stats.lastSharePrice, a.stats.emaSharePrice)
	cost += -normalize(a.stats.lastPB, a.stats.avgPB)

	if a.stats.avgEquityGLA > 0.1 {
		cost += -normalize(a.stats.lastEquityGLA, a.stats.avgEquityGLA)
		div = 3.0
	}

	return cost / div
}

/*============================================================================*
 * Write()                                                                    *
 *============================================================================*/

// Writes an asset to a file.
func (a *Asset) Write(file *os.File) {
	fmt.Fprintf(file, "%s\n", a.ticker)
	fmt.Fprintf(file, "    Last Price %.2f\n", a.stats.lastSharePrice)
	fmt.Fprintf(file, "    Avg. Price %.2f\n", a.stats.avgSharePrice)
	fmt.Fprintf(file, "    AAGR       %.2f\n", a.stats.aagrSharePrice)
	fmt.Fprintf(file, "    Last P/B   %.2f\n", a.stats.lastPB)
	fmt.Fprintf(file, "    Avg. P/B   %.2f\n", a.stats.avgPB)
	fmt.Fprintf(file, "    Avg. DY    %.2f\n", a.stats.avgDY)
	fmt.Fprintf(file, "    EMA  DY    %.2f\n", a.stats.emaDY)
	fmt.Fprintf(file, "\n")
}
