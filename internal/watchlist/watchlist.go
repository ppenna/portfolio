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

package watchlist

import (
	"bufio"
	"fmt"
	"os"
	"portfolio/internal/asset"
	"portfolio/internal/config"
	"portfolio/internal/database"
	"strings"
)

// Watchlist
type Watchlist struct {
	assets []*asset.Asset
}

// Creates an empty watchlist.
func New() *Watchlist {

	w := &Watchlist{}

	w.assets = make([]*asset.Asset, 0)

	return w
}

// Loads an watchlist from a file.
func (w *Watchlist) Load(filename string) error {

	file, err := os.Open(config.WatchlistsPath + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Read watchlist.
	for {

		// Done.
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		ticker := strings.TrimSuffix(line, "\n")
		a, err := database.GetAssetByTicker(ticker)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		w.assets = append(w.assets, a)
	}

	return nil
}

// Returns the list of assets in a watchlist.
func (w *Watchlist) Assets() []*asset.Asset {
	return w.assets
}
