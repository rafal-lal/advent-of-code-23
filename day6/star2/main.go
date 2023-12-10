
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	timeAsStr := strings.Join(strings.Fields(inputLines[0])[1:], "")
	distanceAsStr := strings.Join(strings.Fields(inputLines[1])[1:], "")
	time, _ := strconv.Atoi(timeAsStr)
	distanceMax, _ := strconv.Atoi(distanceAsStr)

	numOfWays := 0
	for i := 1; i < time; i++ {
		distance := (time - i) * i
		if distance > distanceMax {
			numOfWays++
		}
	}

	fmt.Printf("Number of ways: %d\n", numOfWays)
}
