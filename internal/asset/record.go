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
	"time"
)

// Historical Record of an Asset
type AssetRecord struct {
	date            time.Time // Date
	sharePrice      float32   // Price
	bvps            float32   // Book Value per Share
	marketCap       float32   // Market Capitalization
	equity          float32   // Equity
	dividends       float32   // Dividends
	ffo             float32   // Funds from Operations
	numShares       int       // Number of Shares
	defaultRatio    float32   // Default Ratio
	gla             int       // Gross Leasable Area
	numShareHolders int       // Number of Share Holders
}

func readRecord(line []string) *AssetRecord {
	record := &AssetRecord{}

	// Date
	record.date, _ = time.Parse("2006-01-02", line[0])

	// Share Price
	if line[1] != "na" {
		fmt.Sscanf(line[1], "%f", &record.sharePrice)
	}

	// Book Value Price per Share
	if line[2] != "na" {
		fmt.Sscanf(line[2], "%f", &record.bvps)
	}

	// Market Capitalization
	if line[3] != "na" {
		fmt.Sscanf(line[3], "%f", &record.marketCap)
	}

	// Equity
	if line[4] != "na" {
		fmt.Sscanf(line[4], "%f", &record.equity)
	}

	// Dividends
	if line[5] != "na" {
		fmt.Sscanf(line[5], "%f", &record.dividends)
	}

	// Funds from Operations
	if line[6] != "na" {
		fmt.Sscanf(line[6], "%f", &record.ffo)
	}

	// Number of Shares
	if line[7] != "na" {
		fmt.Sscanf(line[7], "%d", &record.numShares)
	}

	// Default Ratio
	if line[8] != "na" {
		fmt.Sscanf(line[8], "%f", &record.defaultRatio)
	}

	// Gross Leasable Area
	if line[9] != "na" {
		fmt.Sscanf(line[9], "%d", &record.gla)
	}

	// Number of Share Holders
	if line[10] != "na" {
		fmt.Sscanf(line[10], "%d", &record.numShareHolders)
	}

	return record
}
