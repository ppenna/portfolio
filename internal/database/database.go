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

package database

import (
	"fmt"
	"portfolio/internal/asset"
	"portfolio/internal/config"
)

// Database Entry
type dbEntry struct {
	ticker string // Ticker
	class  int    // Class
}

// Classes
const (
	Retail = iota
	Mortgage
	Office
	Industrial
	FoF
)

var classesDB = map[int]string{
	Retail:     "Retail",
	Mortgage:   "Mortgage",
	Office:     "Office",
	Industrial: "Industrial",
	FoF:        "FoF",
}

// IDs
const (
	ALZR11 = iota
	BPFF11
	BRCR11
	HGBS11
	HGFF11
	HGLG11
	HGRE11
	JSRE11
	KNCR11
	KNIP11
	VISC11
	XPLG11
	XPML11
)

// Tickers
const (
	TickerALZR11 = "alzr11"
	TickerBPFF11 = "bpff11"
	TickerBRCR11 = "brcr11"
	TickerHGBS11 = "hgbs11"
	TickerHGFF11 = "hgff11"
	TickerHGLG11 = "hglg11"
	TickerHGRE11 = "hgre11"
	TickerJSRE11 = "jsre11"
	TickerKNCR11 = "kncr11"
	TickerKNIP11 = "knip11"
	TickerVISC11 = "visc11"
	TickerXPLG11 = "xplg11"
	TickerXPML11 = "xpml11"
)

// Known Assets
var assetDB = []dbEntry{
	{TickerALZR11, Industrial},
	{TickerBPFF11, FoF},
	{TickerBRCR11, Office},
	{TickerHGBS11, Retail},
	{TickerHGFF11, FoF},
	{TickerHGLG11, Industrial},
	{TickerHGRE11, Office},
	{TickerJSRE11, Office},
	{TickerKNCR11, Mortgage},
	{TickerKNIP11, Mortgage},
	{TickerVISC11, Retail},
	{TickerXPLG11, Industrial},
	{TickerXPML11, Retail},
}

// Assets database.
var database []*asset.Asset

// Loads storage.
func Load() {

	// Nothing to do.
	if database != nil {
		return
	}

	database = make([]*asset.Asset, 0)

	// Load database.
	fmt.Println("Loading database...")
	for i := range assetDB {
		filename := config.DataPath + "/" + assetDB[i].ticker + ".csv"
		a := asset.Read(i, assetDB[i].ticker, filename, assetDB[i].class)
		database = append(database, a)
	}
}

// Returns the list of known assets.
func Assets() []*asset.Asset {
	return database
}

// Gets the ticker of an asset
func AssetTicker(assetID int) (string, error) {

	return assetDB[assetID].ticker, nil
}

// Get asset ID.
func GetAssetID(ticker string) (int, error) {

	// Look for asset.
	for i := range database {
		// Found.
		if database[i].Ticker() == ticker {
			return database[i].ID(), nil
		}
	}

	return -1, fmt.Errorf("unkown ticker " + ticker)

}

// Get asset by ID.
func GetAssetByID(ID int) (*asset.Asset, error) {
	return database[ID], nil
}

// Get asset by ticker.
func GetAssetByTicker(ticker string) (*asset.Asset, error) {

	// Look for asset.
	for i := range database {
		// Found.
		if database[i].Ticker() == ticker {
			return database[i], nil
		}
	}

	return nil, fmt.Errorf("unkown ticker " + ticker)
}

func GetClassName(classID int) (string, error) {
	className, _ := classesDB[classID]

	return className, nil
}
