package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	gameIdsSum := 0

Outer:
	for i, line := range inputLines {
		line = strings.Replace(line, ",", " ", -1)
		line = strings.Replace(line, ":", " ", -1)
		line += ";"
		fields := strings.Fields(line)

		gameId := i + 1
		colorCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		lastVal := 0
		for _, field := range fields[2:] {
			fieldAsInt, err := strconv.Atoi(field)
			if err == nil {
				lastVal = fieldAsInt
			} else if strings.Contains(field, ";") {
				colorCount[field[:len(field)-1]] += lastVal

				if colorCount["red"] > 12 || colorCount["blue"] > 14 || colorCount["green"] > 13 {
					continue Outer
				} else {
					colorCount["red"] = 0
					colorCount["green"] = 0
					colorCount["blue"] = 0
				}
			} else {
				colorCount[field] += lastVal
			}
		}

		gameIdsSum += gameId
	}

	fmt.Printf("Game IDs sum: %d\n", gameIdsSum)
}
