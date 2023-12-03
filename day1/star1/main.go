package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")
	
	var (
		firstDigit int = -99
		lastDigit int
		digit int
		numAsString string
		numAsInt int
	)

	numbersOfEachLine := make([]int, 0)
	for _, line := range inputLines {
		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}
			digit, _ = strconv.Atoi(string(char))

			if firstDigit == -99 {
				firstDigit = digit
			}
			lastDigit = digit
		}
		numAsString = string([]byte{
			byte(strconv.FormatInt(int64(firstDigit), 10)[0]),
			byte(strconv.FormatInt(int64(lastDigit), 10)[0]),
		})
		numAsInt, _ = strconv.Atoi(numAsString)

		numbersOfEachLine = append(numbersOfEachLine, numAsInt)
		firstDigit = -99
	}

	fmt.Printf("Coordinates: %d\n", utils.SumSlice[int](numbersOfEachLine))
}