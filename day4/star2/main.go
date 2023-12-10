package main

import (
	"fmt"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	var copiesSum int
	for i := range inputLines {
		numCopies := countHits(i, inputLines)
		copiesSum += numCopies + 1
	}

	fmt.Printf("Total scratchcards: %d\n", copiesSum)
}

func countHits(i int, lines []string) int {
	splittedLine := strings.Split(lines[i], "|")

	winningNumbersStr := splittedLine[0]
	drawnNumbersStr := splittedLine[1]

	winningNumbers := strings.Fields(winningNumbersStr)[2:]
	drawnNumbers := strings.Fields(drawnNumbersStr)

	hits := 0
	for _, winNum := range winningNumbers {
		for _, drawnNum := range drawnNumbers {
			if winNum == drawnNum {
				hits++
				break
			}
		}
	}

	copiesSum := hits
	// fmt.Printf("line: %d, hits: %d\n", i + 1, hits)
	if hits >= 1 {
		for j := i+1; j < i+1+hits; j++ {
			if j == len(lines) - 1 {
				break
			}
			copiesSum += countHits(j, lines)
		}
	}
	
	return copiesSum
}
