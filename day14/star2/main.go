package main

import (
	"fmt"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	pattern := make([][]string, 0)
	for i, line := range inputLines {
		pattern = append(pattern, []string{})
		for _, ch := range line {
			pattern[i] = append(pattern[i], string(ch))
		}
	}

	states := make([][][]string, 0, 10000)
	// cycles
	var occur1 int
	var clone int
outmost:
	for c := range 1_000_000_000 {
		// north
		swap := true
		for swap {
			swap = false
			for j := len(pattern[0]) - 1; j >= 0; j-- {
				for i := 1; i < len(pattern); i++ {
					if pattern[i][j] == "O" && pattern[i-1][j] == "." {
						pattern[i][j] = "."
						pattern[i-1][j] = "O"
						swap = true
					}
				}
			}
		}
		// west
		swap = true
		for swap {
			swap = false
			for i := 0; i < len(pattern); i++ {
				for j := len(pattern[0]) - 1; j > 0; j-- {
					if pattern[i][j] == "O" && pattern[i][j-1] == "." {
						pattern[i][j] = "."
						pattern[i][j-1] = "O"
						swap = true
					}
				}
			}
		}
		// south
		swap = true
		for swap {
			swap = false
			for j := 0; j < len(pattern[0]); j++ {
				for i := 0; i < len(pattern)-1; i++ {
					if pattern[i][j] == "O" && pattern[i+1][j] == "." {
						pattern[i][j] = "."
						pattern[i+1][j] = "O"
						swap = true
					}
				}
			}
		}
		// east
		swap = true
		for swap {
			swap = false
			for i := 0; i < len(pattern); i++ {
				for j := 0; j < len(pattern[0])-1; j++ {
					if pattern[i][j] == "O" && pattern[i][j+1] == "." {
						pattern[i][j] = "."
						pattern[i][j+1] = "O"
						swap = true
					}
				}
			}
		}

		sum := 0
		cnt := len(pattern)
		for i := range pattern {
			for j := range pattern[i] {
				if pattern[i][j] == "O" {
					sum += cnt
				}
			}
			cnt--
		}

		for i, state := range states {
			same := true
			for i := range pattern {
				for j := range pattern[0] {
					if pattern[i][j] != state[i][j] {
						same = false
						break
					}
				}
			}
			if same {
				// printPattern(pattern)
				println(i)
				occur1 = i
				println(c)
				clone = c
				println(sum)
				break outmost
			}
		}

		patternCopy := make([][]string, len(pattern))
		for i := range pattern {
			patternCopy[i] = make([]string, len(pattern[0]))
			copy(patternCopy[i], pattern[i])
		}
		states = append(states, patternCopy)
	}

	// every 84 cycles starting from 94th pattern looks the same
	shifts := (1_000_000_000 - (occur1+1)) % (clone - occur1)
	finalPattern := states[occur1+shifts]

	sum := 0
	cnt := len(finalPattern)
	for i := range finalPattern {
		for j := range finalPattern[i] {
			if finalPattern[i][j] == "O" {
				sum += cnt
			}
		}
		cnt--
	}

	// printPattern(finalPattern)
	println(sum)
}

func printPattern(pattern [][]string) {
	for i := range pattern {
		for j := range pattern[i] {
			fmt.Printf("%s", pattern[i][j])
		}
		println()
	}
}
