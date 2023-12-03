package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	digits := []string{
		"one",   "two",   "three",
		"four",  "five",  "six",
		"seven", "eight", "nine",
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9",
	}

	var (
		lowestIndex  int = 10_0000
		highestIndex int = -1
		lowestDigit  string
		highestDigit string
	)

	numbersOfEachLine := make([]int, 0)
	for _, line := range inputLines {
		for _, digit := range digits {
			if strings.Contains(line, digit) {
				idxFirst := strings.Index(line, digit)
				if idxFirst < lowestIndex {
					lowestDigit = digit
					lowestIndex = idxFirst
				}

				idxLast := strings.LastIndex(line, digit) + len(digit)
				if idxLast > highestIndex {
					highestDigit = digit
					highestIndex = idxLast
				}
			}
		}

		if len(lowestDigit) != 1 {
			lowestDigit = numberMap[lowestDigit]
		}
		if len(highestDigit) != 1 {
			highestDigit = numberMap[highestDigit]
		}

		numAsInt, _ := strconv.Atoi(lowestDigit + highestDigit)
		numbersOfEachLine = append(numbersOfEachLine, numAsInt)
		lowestIndex = 10_0000
		highestIndex = -1
	}

	fmt.Printf("Coordinates %d\n", utils.SumSlice[int](numbersOfEachLine))
}
