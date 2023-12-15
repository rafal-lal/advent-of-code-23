package main

import (
	"fmt"
	//"math"
	"strconv"
	"strings"
	"slices"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	reports := make([][]int, len(inputLines))
	for i, line := range inputLines {
		fields := strings.Fields(line)
		reports[i] = make([]int, len(fields))
		for j, field := range fields {
			num, _ := strconv.Atoi(field)
			reports[i][j] = num
		}
	}

	sumExtrapol := 0
	for _, report := range reports {
		rightVals := make([]int, 0)
		currentDiffs := report

		for {
			newDiffs := make([]int, 0)
			for i := 0; i < len(currentDiffs)-1; i++ {
				//newDiffs = append(newDiffs, int(math.Abs(float64(currentDiffs[i+1])-float64(currentDiffs[i]))))
				newDiffs = append(newDiffs, currentDiffs[i+1]-currentDiffs[i])
			}
			rightVals = append(rightVals, newDiffs[len(newDiffs)-1])

			allZeros := true
			for _, val := range newDiffs {
				if val != 0 {
					allZeros = false
					break
				}
			}

			if !allZeros {
				currentDiffs = newDiffs
				continue
			}

			slices.Reverse(rightVals)

			lastVal := 0
			for i := 1; i < len(rightVals); i++ {
				lastVal = lastVal + rightVals[i]
			}

			sumExtrapol += lastVal + report[len(report)-1]

			break
		}
	}

	fmt.Printf("Sum of extrapolated values is %d\n", sumExtrapol)
}
