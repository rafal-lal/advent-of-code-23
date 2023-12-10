package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	var races [4][2]int
	times := strings.Fields(inputLines[0])[1:]
	distances := strings.Fields(inputLines[1])[1:]

	for i := range races {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		races[i][0] = time
		races[i][1] = distance
	}

	numOfWaysProduct := 1
	for _, race := range races {
		numOfWays := 0
		for i := 1; i < race[0]; i++ {
			distance := (race[0] - i) * i
			if distance > race[1] {
				numOfWays++
			}
		}
		numOfWaysProduct *= numOfWays
	}

	fmt.Printf("Number of ways: %d\n", numOfWaysProduct)
}
