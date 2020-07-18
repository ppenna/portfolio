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
	"portfolio/internal/utils"
)

// Asset Statistics
type AssetStatistics struct {
	lastSharePrice float32 // Last Share Price
	avgSharePrice  float32 // Average Share Price
	emaSharePrice  float32 // EMA Share Price
	aagrSharePrice float32 // Averave Annual Growth Rate

	// Equity/GLA Statistics
	lastEquityGLA float32
	avgEquityGLA  float32

	// P/B Statistics
	lastPB float32 // Last P/B
	avgPB  float32 // Average P/B

	// DY Statistics
	avgDY float32 // Average Dividend Yield
	emaDY float32 // EMA Dividend Yield
}

// Compute the EMA of a serie.
func computeEMA(x []float32, t, n int) float32 {
	var ema func([]float32, int) float32

	k := float32(2.0) / float32(n+1)

	ema = func(x []float32, t int) float32 {
		if t == 0 {
			return 0.0
		}

		return x[t]*k + ema(x, t-1)*(1.0-k)
	}

	return ema(x, t)
}

// Computes statistics on dividend yield.
func (stats *AssetStatistics) computeDY(hist *AssetHistory) {

	histDY := make([]float32, 0)

	// Compute average statistics.
	for _, record := range hist.records {
		dy := 12.0 * record.dividends / record.sharePrice

		stats.avgDY += dy
		histDY = append(histDY, dy)
	}

	stats.emaDY = computeEMA(histDY, len(histDY)-1, 6)
	stats.avgDY /= float32(len(hist.records))
}

// Computes statistics on share price.
func (stats *AssetStatistics) computeSharePrice(hist *AssetHistory) {

	// First and last records.
	firstRecord := hist.records[0]
	lastRecord := hist.records[len(hist.records)-1]

	// Last Share Price
	stats.lastSharePrice = lastRecord.sharePrice

	// Last P/B
	stats.lastPB = lastRecord.sharePrice / lastRecord.bvps

	// Last Equity GLA
	stats.lastEquityGLA = lastRecord.equity / float32(lastRecord.gla)

	// Compute average statistics.
	histSharePrice := make([]float32, 0)
	for _, record := range hist.records {
		stats.avgSharePrice += record.sharePrice
		stats.avgPB += record.sharePrice / record.bvps

		if record.gla > 0 {
			stats.avgEquityGLA += record.equity / float32(record.gla)
		}

		histSharePrice = append(histSharePrice, record.sharePrice)
	}
	stats.avgSharePrice /= float32(len(hist.records))
	stats.avgPB /= float32(len(hist.records))
	stats.avgEquityGLA /= float32(len(hist.records))
	stats.emaSharePrice = computeEMA(histSharePrice, len(histSharePrice)-1, 6)

	// Compute average anual growth rate.
	duration := utils.MonthDiff(hist.startDate, hist.endDate)
	priceDevelopment := lastRecord.sharePrice/firstRecord.sharePrice - 1
	stats.aagrSharePrice = 12 * priceDevelopment / float32(duration)
}

// Compute statistics on historical data.
func computeStatistics(hist *AssetHistory) *AssetStatistics {

	stats := &AssetStatistics{}

	stats.computeSharePrice(hist)
	stats.computeDY(hist)

	return stats
}
