package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/1
	// aoc.Local(part1, "part1", "sample.txt", 14)
	aoc.Local(part1, "part1", "input.txt", 74)
	// aoc.Local(part2, "part2", "sample.txt", 34)
	aoc.Local(part2, "part2", "input.txt", 1795)
}

func part1(contents string) interface{} {
	floor := 0
	for _, ch := range contents {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
	}
	return floor
}

func part2(contents string) interface{} {
	floor := 0
	for i, ch := range contents {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
