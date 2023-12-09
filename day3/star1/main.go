package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	// Create 2D array from lines
	arr := make([][]rune, len(inputLines))
	for i := range arr {
		arr[i] = make([]rune, len(inputLines[i]))
		for j := range arr[i] {
			arr[i][j] = rune(inputLines[i][j])
		}
	}

	var sum int
	for i := range arr {
		var currentNumber string
		lastWasNumber := false
		isCurrentNumberAdjacent := false
		for j := range arr[i] {
			if !unicode.IsNumber(arr[i][j]) {
				if isCurrentNumberAdjacent {
					num, _ := strconv.Atoi(currentNumber)
					println(num)
					sum += num
					isCurrentNumberAdjacent = false
				}
				lastWasNumber = false
				continue
			}

			if lastWasNumber {
				currentNumber += string(arr[i][j])
			} else {
				currentNumber = string(arr[i][j])
				lastWasNumber = true
			}

			if !isCurrentNumberAdjacent {
				isCurrentNumberAdjacent = isNumberAdjacent(i, j, arr)
			}
		}
		if isCurrentNumberAdjacent {
			num, _ := strconv.Atoi(currentNumber)
			println(num)
			sum += num
			isCurrentNumberAdjacent = false
		}
	}
	
	fmt.Printf("Engine schematic sum: %d\n", sum)
}

func isNumberAdjacent(i, j int, arr [][]rune) bool {
	if i != 0 {
		// check N
		if !unicode.IsDigit(arr[i-1][j]) && arr[i-1][j] != '.' {
			return true
		}

		if j != 0 {
			// check NW, W
			if (!unicode.IsDigit(arr[i-1][j-1]) && arr[i-1][j-1] != '.') || (!unicode.IsDigit(arr[i][j-1]) && arr[i][j-1] != '.') {
				return true
			}
		}

		if j != len(arr[0]) - 1 {
			// check NE, E
			if (!unicode.IsDigit(arr[i-1][j+1]) && arr[i-1][j+1] != '.') || (!unicode.IsDigit(arr[i][j+1]) && arr[i][j+1] != '.') {
				return true
			}
		}
	}

	if i != len(arr) - 1 {
		// check S
		if !unicode.IsDigit(arr[i+1][j]) && arr[i+1][j] != '.' {
			return true
		}

		if j != 0 {
			// check SW, W
			if (!unicode.IsDigit(arr[i+1][j-1]) && arr[i+1][j-1] != '.') || (!unicode.IsDigit(arr[i][j-1]) && arr[i][j-1] != '.') {
				return true
			}
		}

		if j != len(arr[0]) - 1 {
			// check SE, E
			if (!unicode.IsDigit(arr[i+1][j+1]) && arr[i+1][j+1] != '.') || (!unicode.IsDigit(arr[i][j+1]) && arr[i][j+1] != '.') {
				return true
			}
		}
	}

	return false
}
