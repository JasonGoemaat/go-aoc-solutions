package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	aoc.Local(part1, "part1", "sample.txt", 1)
	aoc.Local(part1, "part1", "sample2.txt", 36)
	aoc.Local(part1, "part1", "input.txt", 682)
	aoc.Local(part2, "part2", "sample3.txt", 3)
	aoc.Local(part2, "part2", "sample4.txt", 13)
	aoc.Local(part2, "part2", "sample5.txt", 227)
	aoc.Local(part2, "part2", "input.txt", 1511)
}

func countNines(a *aoc.Area, v map[int]bool, r, c int, h byte) int {
	if !a.Inside(r, c) {
		return 0
	}
	if a.Get(r, c) != h {
		return 0
	}
	pos := a.RowColToIndex(r, c)
	if v[pos] {
		return 0
	}
	v[pos] = true
	if h == 0x39 {
		return 1
	}
	total := countNines(a, v, r+1, c, h+1)
	total += countNines(a, v, r-1, c, h+1)
	total += countNines(a, v, r, c+1, h+1)
	total += countNines(a, v, r, c-1, h+1)
	return total
}

func part1(contents string) interface{} {
	area := aoc.ParseArea(contents)
	sum := 0

	// need to find each 0, then search for 9s, avoid double-counting
	for p := 0; p < len(area.Data); p++ {
		if area.GetIndex(p) == 0x30 {
			visited := map[int]bool{}
			r, c := area.IndexToRowCol(p)
			sum += countNines(area, visited, r, c, byte(0x30))
		}
	}
	return sum
}

// WIP using maps
func part2a(contents string) interface{} {
	area := aoc.ParseArea(contents)
	sum := 0

	// need to find each 0, then search for 9s, avoid double-counting
	for p := 0; p < len(area.Data); p++ {
		if area.GetIndex(p) == 0x30 {
			list := map[int]int{0x30: p}
			// newList := map[int]int{}
			h := 0x30
			for ; h <= 39; h++ {
				for p2, count := range list {
					delete(list, p2)
					r, c := area.IndexToRowCol(p2)
					if area.Inside(r-1, c) && area.Get(r-1, c) == byte(h+1) {
						p3 := area.RowColToIndex(r-1, c)
						existing, ok := list[p3]
						if ok {
							list[p3] = existing + count
						} else {
							list[p3] = count
						}
					}
				}
			}
		}
	}
	return sum
}

// trying another way
func part2(contents string) interface{} {
	area := aoc.ParseArea(contents)
	counts := make([]int, len(area.Data))

	totalAround := func(position int) int {
		h := area.GetIndex(position)
		r, c := area.IndexToRowCol(position)
		total := 0
		for _, v := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if area.Inside(r+v[0], c+v[1]) && (area.Get(r+v[0], c+v[1]) == (h - 1)) {
				total += counts[area.RowColToIndex(r+v[0], c+v[1])]
			}
		}
		return total
	}

	// initialize '0' location to count as 1, when we look for '0' aroung '1',
	// this will count each one as a path
	for p := range area.Data {
		if area.GetIndex(p) == 0x30 {
			counts[p] = 1
		}
	}

	for h := byte(0x31); h <= 0x39; h++ {
		for p := range area.Data {
			if area.GetIndex(p) == h {
				counts[p] = totalAround(p)
			}
		}
	}

	// now we count the totals for '9's
	sum := 0
	for p := range area.Data {
		if area.GetIndex(p) == 0x39 {
			sum += counts[p]
		}
	}
	return sum
}
