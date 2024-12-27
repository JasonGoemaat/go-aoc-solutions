package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.Local(part1, "part1", "sample.aoc", 0)
	aoc.Local(part1, "part1", "input.aoc", 0)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0)
}

func incrementPassword(password []byte, index int) {
	password[index]++
	if password[index] > 'z' {
		incrementPassword(password, index-1)
	} else {
		for i := index + 1; i < len(password); i++ {
			password[i] = 'a'
		}
	}
}

// handle invalid characters 'i', 'o', 'l' by incrementing the right-most
// and changing all characters after it to 'a'.   Return true if modified
func handleInvalidPassword(password []byte) bool {
	for i := range len(password) {
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			incrementPassword(password, i)
			return true // modified
		}
	}
	return false // unmodified
}

func goodPassword(password []byte) bool {
	// scan for incrementing 3 bytes
	doubleIndex1, doubleIndex2 := -1, -1
	straightIndex := -1
	for i := 1; i < len(password); i++ {
		if i > 1 && straightIndex == -1 {
			if password[i-1] == (password[i]-1) && password[i-2] == (password[i]-2) {
				straightIndex = i
				continue
			}
		}
		if password[i] == password[i-1] {
			if doubleIndex1 == -1 {
				doubleIndex1 = i
				continue
			}
			if doubleIndex2 == -1 && doubleIndex1 < (i-1) {
				doubleIndex1 = i
				continue
			}
		}
	}
	// ways to speed it up:
	// 1. some way to increment as far right as posssible far enough
	//    to make something work.   I.e. if we have first 3 characters forming
	//	  a straight, and next 2 forming a pair, we need to increment last
	//    until it gets a pair, or last 2 will become 'aa' to make the pair.
	// 2. Likewise if we have two pairs early, we increment last 3 to get to straight
	// 3. Therefore, we split string into 'work' bytes and set straights and pairs.
	// 4. And if we can merge last characters to make the final thing, do it...
	// SCRATCH THAT
	// Ok, the tricky part is handling possible matches earlier.  It would be nice
	// if we could start with the assumption that nothing matched or would match
	// anything.   Maybe we can?   But there are two things we have to match or it
	// would be super easy.   In most cases we need the right-most 7 characters to
	// fit the rules.   The only EDGE CASE is if the first 2 characters can fit the
	// rules.  So when I see a string ending in 'gk', I know I need to increment
	// the 'g' because  'k-z' for the last character doesn't fit either of the tropes.
	// Actually then I start with 'hh' because that fits at least one rule.   I can
	// increment to 'hi' if the character before is 'g', but otherwise it won't
	// matter.  And if there aren't any earlier, I should move through the whole
	// string incrementing until something fits.
	// So what to do first?  Oh, hold up, this is the clue that tells you all you need
	// to know:
	//		The next password after ghijklmn is ghjaabcc
	// So the triple run CAN be overlapping with the two doubles and has to be in the
	// middle if it is.  Maybe I can figure out my input 'hepxcrrq' manually.  Only
	// the first 5 characters are required.   I can change the end to 'xxyzz' without
	// having to do anything to the first 'hepx'.  YEP!  That worked! 'hepxxyzz'
	// PART 2: ok, it expired again.   I can't use the same trick and since the last
	// digits are 'z', we need to roll-over the 'x'.   Since we only need the first 5
	// to form the lowest 'aabcc', we just roll the 'p' up to 'q' and use 'aabcc' for
	// 'heqaabcc'
	return doubleIndex1 > 0 && doubleIndex2 > 0 && straightIndex > 0
}

// simple, 1 input, first output
func part1(contents string) interface{} {
	bytes := []byte(contents)
	for true {
		if handleInvalidPassword(bytes) {
			continue
		}
		if goodPassword(bytes) {
			break
		}
		incrementPassword(bytes, len(bytes)-1)
	}
	return string(bytes)
}

func part2(contents string) interface{} {
	return 0
}
