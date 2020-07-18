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
	"math/rand"
	"sort"
)

// GA Configuration
const (
	populationSize  = 5000  // Population Size
	selectionRatio  = 0.60  // Selection Ratio
	eliteRatio      = 0.05  // Elite Ratio
	mutationRatio   = 0.02  // Mutation Ratio
	evolutionCutOff = 1000  // GA Evolution Cut Off
	maxGenerations  = 10000 // Maximum Number of Generations
)

/*============================================================================*
 * Gene                                                                       *
 *============================================================================*/

// Gene
type gene struct {
	dna     []float32
	fitness float32 // Fitness
}

// Normalizes a gene.
func (g *gene) normalize() {
	norm := float32(0.0)

	for i := 0; i < len(g.dna); i++ {
		norm += g.dna[i]
	}

	if norm > 0.0 {
		for i := 0; i < len(g.dna); i++ {
			g.dna[i] /= norm
		}
	}
}

// Creates a new gene.
func newGene() *gene {
	g := &gene{}

	g.dna = NewAllocation()

	g.normalize()

	return g
}

// Evaluates the fitness of a gene.
func (g *gene) eval() {
	g.fitness = eval(g.dna)
}

// Cross overs two genes.
func crossover(g1, g2 *gene) *gene {

	g := &gene{}
	g.dna = make([]float32, len(assets))

	point := len(g.dna) / 2

	for i := 0; i < point; i++ {
		g.dna[i] = g2.dna[i]
	}

	for i := point; i < len(g.dna); i++ {
		g.dna[i] = g2.dna[i]
	}

	for i := 0; i < len(g.dna); i++ {
		if g.dna[i] < minAllocation {
			g.dna[i] = 0.0
		}
	}

	g.normalize()

	return g
}

// Mutates a gene.
func (g *gene) mutate() {
	point := rand.Int31n(int32(len(g.dna)))

	g.dna[point] = rand.Float32()

	g.normalize()
}

/*============================================================================*
 * Genetic Algorithm                                                          *
 *============================================================================*/

// Genetic Algorithms
type GeneticAlgorithm struct {
	populationSize int // Population Size
	selectionSize  int // Selection Ratio
	eliteSize      int // Elite Ratio
}

// Instantiates a new genetic algorithm
func newGeneticAlgorithm(popSize int, sRatio, eRatio, mRatio float32) *GeneticAlgorithm {
	ga := &GeneticAlgorithm{}

	ga.populationSize = popSize
	ga.selectionSize = int(sRatio * float32(ga.populationSize))
	ga.eliteSize = int(eRatio * float32(ga.populationSize))

	return ga
}

// Selects organisms to mate.
func (ga *GeneticAlgorithm) selection(population []*gene) []*gene {

	parents := make([]*gene, 0)

	totalFitness := float32(0.0)
	for _, g := range population {
		totalFitness += g.fitness
	}

	for i := 0; i < ga.selectionSize; i++ {
		done := false
		f := rand.Float32() * totalFitness

		for !done {
			for _, g := range population {
				f -= g.fitness

				// Found.
				if f <= 0 {
					parents = append(parents, g)
					done = true
					break
				}
			}
			done = true
		}
	}

	return parents
}

// Breed new genes.
func (ga *GeneticAlgorithm) breed(parents []*gene) []*gene {
	children := make([]*gene, 0)

	for i := 0; i < len(parents)-2; i += 2 {
		child := crossover(parents[i], parents[i+1])

		children = append(children, child)
	}

	for _, g := range children {

		if rand.Float32() <= mutationRatio {
			g.mutate()
		}

		g.eval()
	}

	return children
}

// Replace old population.
func (ga *GeneticAlgorithm) replace(population, children []*gene) {
	for _, child := range children {
		i := rand.Int31n(int32(ga.populationSize - ga.eliteSize))

		population[i] = child
	}
}

type ByFitness []*gene

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Less(i, j int) bool { return a[i].fitness < a[j].fitness }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Runs the target genetic algorithm
func (ga *GeneticAlgorithm) run(verbose bool) *gene {

	var bestGene *gene

	genes := make([]*gene, ga.populationSize)

	// Generate initial population.
	for i := 0; i < ga.populationSize; i++ {
		genes[i] = newGene()
		genes[i].eval()
	}
	sort.Sort(ByFitness(genes))

	bestGene = genes[len(genes)-1]

	if verbose {
		fmt.Println("Running Genetic Algorithm...")
	}

	lastGeneration := evolutionCutOff
	for i := 1; i < maxGenerations; i++ {
		parents := ga.selection(genes)

		children := ga.breed(parents)
		ga.replace(genes, children)
		sort.Sort(ByFitness(genes))

		if genes[len(genes)-1].fitness > bestGene.fitness {
			bestGene = genes[len(genes)-1]

			lastGeneration = i + evolutionCutOff

			if verbose {
				fmt.Printf("%4d Best Fitness: %f\n", i, bestGene.fitness)

			}
		}

		// Stable solution.
		if i >= lastGeneration {
			break
		}
	}

	return bestGene
}
