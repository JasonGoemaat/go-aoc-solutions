package year2015day15part2

import (
	"fmt"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func calculateScore() int {
	for i := range len(ingredients[0]) {
		scores[i] = 0
	}

	for i, ingredient := range ingredients {
		for j := range len(ingredient) {
			score := ingredient[j] * counts[i]
			scores[j] += score
		}
	}

	// must have precisely 500 calories
	if scores[len(scores)-1] != 500 {
		return -1000000
	}

	total := scores[0]
	for i := 1; i < len(scores)-1; i++ {
		total *= max(scores[i], 0)
	}
	return total
}

var ingredients [][]int
var counts []int
var scores []int
var totalRecurseCalls = 0
var finalCalls = 0

func recurse(index, remaining int) int {
	totalRecurseCalls++
	if index == (len(counts) - 1) {
		counts[index] = remaining
		finalCalls++
		return calculateScore()
	}
	best := 0
	for i := remaining; i >= 0; i-- {
		counts[index] = i
		best = max(best, recurse(index+1, remaining-i))
	}
	return best
}

func Part2(contents string) interface{} {
	// bah, this is tough
	ingredients = aoc.ParseIntsPerLine(contents) // name doesn't matter
	counts = make([]int, len(ingredients))
	scores = make([]int, len(ingredients[0]))

	// will have to try 94 million options
	bestScore := recurse(0, 100)
	fmt.Printf("Total recurse() calls: %d\n", totalRecurseCalls)
	fmt.Printf("Total recurse() calls at the end: %d\n", finalCalls)
	return bestScore
}
