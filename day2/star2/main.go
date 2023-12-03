package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	sumOfProducts := 0
	for _, line := range inputLines {
		line = strings.Replace(line, ",", " ", -1)
		line = strings.Replace(line, ":", " ", -1)
		line = strings.Replace(line, ";", " ", -1)
		fields := strings.Fields(line)

		colorMaxVal := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		lastVal := 0
		for _, field := range fields[2:] {
			fieldAsInt, err := strconv.Atoi(field)
			if err == nil {
				lastVal = fieldAsInt
			} else {
				if lastVal > colorMaxVal[field] {
					colorMaxVal[field] = lastVal
				}
			}			
		}

		product := 1
		for _, val := range colorMaxVal {
			product *= val
		}
		sumOfProducts += product
	}

	fmt.Printf("Game IDs sum: %d\n", sumOfProducts)
}
