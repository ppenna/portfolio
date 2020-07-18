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

package wallet

import (
	"bufio"
	"fmt"
	"os"
	"portfolio/internal/config"
	"portfolio/internal/database"
	"strings"
)

// Wallet
type Wallet struct {
	name        string          // Name
	allocation  map[int]float32 // Allocation
	performance float32         // Performance
	price       float32         // Cost
}

// Creates an empty wallet.
func New(name string) *Wallet {
	wallet := &Wallet{}

	wallet.name = name
	wallet.allocation = make(map[int]float32)

	return wallet
}

// Instantiates my current wallet.
func MyWallet() (*Wallet, error) {
	myWallet, err := Read("default.wallet")
	if err != nil {
		return nil, err
	}

	return myWallet, nil
}

// Sets an allocation for a wallet.
func (wallet *Wallet) SetAllocation(newAllocation map[int]float32) {

	// TODO check arguments

	wallet.allocation = newAllocation
}

/*============================================================================*
 * Performance()                                                              *
 *============================================================================*/

// Computes the performance of the target wallet.
func (wallet *Wallet) Performance() float32 {
	wallet.performance = 0.0

	for i := range wallet.allocation {
		if wallet.allocation[i] > 0.0 {
			a, _ := database.GetAssetByID(i)
			wallet.performance += a.Performance() * wallet.allocation[i] * 100
		}
	}

	return wallet.performance
}

/*============================================================================*
 * Cost()                                                                    *
 *============================================================================*/

// Computes the price of the target wallet.
func (wallet *Wallet) Cost() float32 {
	wallet.price = 0.0

	for i := range wallet.allocation {
		if wallet.allocation[i] > 0.0 {
			a, _ := database.GetAssetByID(i)
			wallet.price += a.Cost() * wallet.allocation[i] * 100
		}
	}

	return wallet.price
}

/*============================================================================*
 * Risk()                                                                     *
 *============================================================================*/

// Computes the risk of the target wallet.
func (wallet *Wallet) Risk() float32 {
	var risk float32
	classes := []float32{0, 0, 0, 0, 0}

	// Compute asset allocation in each class.
	for i := range wallet.allocation {

		a, _ := database.GetAssetByID(i)
		classes[a.Class()] += wallet.allocation[i]
	}

	for i := range classes {
		if classes[i] > risk {
			risk = classes[i]
		}
	}

	return 10 * (1 - risk)
}

/*============================================================================*
 * Read()                                                                     *
 *============================================================================*/

// Reads a wallet from a file.
func Read(filename string) (*Wallet, error) {

	file, err := os.Open(config.WalletsPath + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	allocation := make(map[int]float32)

	// Read wallet name.
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("corrupted wallet file")
	}
	name := strings.TrimSuffix(line, "\n")

	// Read allocations
	for {
		var a float32
		var assetID int
		var assetTicker string

		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Sscanf(line, "%s %f", &assetTicker, &a)

		assetID, err = database.GetAssetID(assetTicker)
		if err != nil {
			break
		}
		allocation[assetID] = a / 100.0
	}

	// Instantiate wallet.
	wallet := &Wallet{}
	wallet.name = name
	wallet.allocation = allocation

	return wallet, nil
}

/*============================================================================*
 * Persist()                                                                  *
 *============================================================================*/

// Persists the allocation of the target wallet into a file.
func (wallet *Wallet) Persist(filename string) error {

	file, err := os.Create(config.WalletsPath + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n", wallet.name)
	for assetID, assetAllocation := range wallet.allocation {
		if assetAllocation > 0.0 {
			fmt.Fprintf(file,
				"%d %5.2f\n",
				assetID,
				assetAllocation*100.0,
			)
		}
	}

	return nil
}

/*============================================================================*
 * Write()                                                                    *
 *============================================================================*/

// Writes the allocation of the target wallet into a file.
func (wallet *Wallet) Write(file *os.File) error {

	// Invalid file.
	if file == nil {
		return fmt.Errorf("invalid wallet file")
	}

	fmt.Fprintf(file, "%s\n", wallet.name)
	for assetID, assetAllocation := range wallet.allocation {
		if assetAllocation > 0.0 {
			ticker, _ := database.AssetTicker(assetID)
			fmt.Fprintf(file,
				"  %s %5.2f\n",
				ticker,
				assetAllocation*100.0,
			)
		}
	}

	return nil
}

/*============================================================================*
 * PrintStats()                                                               *
 *============================================================================*/

func (wallet *Wallet) PrintStats(file *os.File) {
	classes := []float32{0, 0, 0, 0, 0}

	fmt.Fprintf(file, "\nStatistics for %s\n", wallet.name)

	// Compute asset allocation in each class.
	for i := range wallet.allocation {
		a, _ := database.GetAssetByID(i)
		classes[a.Class()] += wallet.allocation[i]
	}

	for i := range classes {
		className, _ := database.GetClassName(i)

		fmt.Fprintf(file,
			"  %-15s %5.2f %%\n",
			className,
			100*classes[i],
		)
	}

	performance := wallet.Performance()
	cost := wallet.Cost()
	risk := wallet.Risk()
	score := (performance + cost + risk) / 3.0

	fmt.Fprintf(file, "\n  %-15s %5.2f %%\n", "Performance", performance)
	fmt.Fprintf(file, "  %-15s %5.2f %%\n", "Cost", cost)
	fmt.Fprintf(file, "  %-15s %5.2f %%\n", "Risk", risk)
	fmt.Fprintf(file, "  %-15s %5.2f %%\n\n", "Score", score)
}
