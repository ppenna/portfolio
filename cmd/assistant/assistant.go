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
	"math/rand"
	"portfolio/internal/asset"
	"portfolio/internal/wallet"
	"portfolio/internal/watchlist"
	"time"
)

// Assistant Configuration
const (
	minAllocation = 0.02 // Minimum Allocation for an Asset
	maxAllocation = 0.15 // Maximum Allocation for an Asset
)

// Assets.
var assets []*asset.Asset

// Generates a random allocation.
func NewAllocation() []float32 {
	allocation := make([]float32, len(assets))

	for i := 0; i < len(assets); i++ {
		allocation[i] = rand.Float32()
	}

	return allocation
}

// Computes the performance valuation of an allocation.
func costEval(allocation []float32) float32 {
	value := float32(0.0)

	for i := range allocation {
		value += assets[i].Cost() * allocation[i]
	}

	return value
}

// Computes the performance valuation of an allocation.
func perfEval(allocation []float32) float32 {
	performance := float32(0.0)

	for i := range allocation {
		performance += assets[i].Performance() * allocation[i]
	}

	return performance
}

func eval(allocation []float32) float32 {
	cost := costEval(allocation)
	performance := perfEval(allocation)
	risk := riskEval(allocation)

	return (cost + performance + risk) / 3.0
}

// Computes the risk valuation of an allocation.
func riskEval(allocation []float32) float32 {
	var risk float32
	classes := []float32{0, 0, 0, 0, 0}

	// Compute asset allocation in each class.
	for i := range allocation {

		// Too risky.
		if allocation[i] >= maxAllocation {
			return 0.0
		}

		classes[assets[i].Class()] += allocation[i]
	}

	for i := range classes {
		if classes[i] > risk {
			risk = classes[i]
		}
	}

	return (1 - risk) / 10.0
}

func AssistantRun(watchlist *watchlist.Watchlist, verbose bool) *wallet.Wallet {

	assets = watchlist.Assets()
	wallet := wallet.New("Recommended Wallet")

	rand.Seed(time.Hour.Nanoseconds())

	assistant := newGeneticAlgorithm(populationSize,
		selectionRatio,
		eliteRatio,
		mutationRatio,
	)

	bestSolution := assistant.run(verbose)
	allocation := make(map[int]float32)
	for i, a := range watchlist.Assets() {
		assetID := a.ID()
		allocation[assetID] = bestSolution.dna[i]
	}

	wallet.SetAllocation(allocation)

	return wallet
}
