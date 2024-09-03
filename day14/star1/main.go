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

	for i := range pattern {
		for j := range pattern[i] {
			fmt.Printf("%s", pattern[i][j])
		}
		println()
	}

	println(sum)
}
