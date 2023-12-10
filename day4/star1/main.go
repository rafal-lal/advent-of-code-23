package main

import (
	"fmt"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	var pointsSum int
	for _, line := range inputLines {
		splittedLine := strings.Split(line, "|")

		winningNumbersStr := splittedLine[0]
		drawnNumbersStr := splittedLine[1]

		winningNumbers := strings.Fields(winningNumbersStr)[2:]
		drawnNumbers := strings.Fields(drawnNumbersStr)

		hits := 0
		for _, winNum := range winningNumbers {
			for _, drawnNum := range drawnNumbers {
				if winNum == drawnNum {
					hits++
					//break
				}
			}
		}

		points := 0
		if hits >= 1 {
			points = 1
			for i := 1; i < hits; i++ {
				points *= 2
			}
		}

		pointsSum += points
	}

	fmt.Printf("Cards value: %d\n", pointsSum)
}
