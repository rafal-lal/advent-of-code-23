package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	var seeds []int
	data := make([][][]int, 0)

	// prepare data
	seedsAsStr := strings.Fields(inputLines[0])
	for _, seed := range seedsAsStr[1:] {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}

	sectionCounter := 0
	lineCounter := 0
	data = append(data, [][]int{})
	for _, line := range inputLines[3:] {
		if line == "" {
			continue
		}
		lineHasNum := regexp.MustCompile(`\d`).MatchString(line)
		if !lineHasNum {
			sectionCounter++
			data = append(data, [][]int{})
			lineCounter = 0
			continue
		}
		
		data[sectionCounter] = append(data[sectionCounter], []int{})
		numsAsStr := strings.Fields(line)
		for _, numAsStr := range numsAsStr {
			num, _ := strconv.Atoi(numAsStr)
			data[sectionCounter][lineCounter] = append(data[sectionCounter][lineCounter], num)
		}
		lineCounter++
	}

	lowestLocation := 2147483647 // highest int
	// match data
	for _, seed := range seeds {
		lastVal := seed
		for _, outerVal := range data {
			currentVal := -1
			for _, val := range outerVal {
				if lastVal >= val[1] && lastVal <= (val[1]+val[2]) {
					currentVal = lastVal + (val[0] - val[1])
				}
			}
			if currentVal == -1 {
				currentVal = lastVal
			}
			lastVal = currentVal
		}

		location := lastVal
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Printf("Lowest location: %d\n", lowestLocation)
}
