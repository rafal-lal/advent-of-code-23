package main

import (
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	steps := strings.Split(inputLines[0], ",")
	sum := 0
	for _, step := range steps {
		localVal := 0
		for _, ch := range step {
			localVal += int(ch)
			localVal *= 17
			localVal %= 256
		}
		sum += localVal
	}

	println(sum)
}
