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
		for j := range arr[i] {
			if arr[i][j] == '*' {
				num := isGear(i, j, arr)
				sum += isGear(i, j, arr)
				println(i, j, num)
				println("")
			}
		}
	}

	fmt.Printf("Gear ratios sum: %d\n", sum)
}

func isGear(i, j int, arr [][]rune) int {
	coordinates := map[string]bool{
		"NW": false, // NW
		"N":  false, // N
		"NE": false, // NE
		"W":  false, // W
		"E":  false, // E
		"SW": false, // SW
		"S":  false, // S
		"SE": false, // SE
	}

	numbers := make([]int, 0)
	if i != 0 {
		// check N
		if unicode.IsDigit(arr[i-1][j]) {
			coordinates["N"] = true
			number := string(arr[i-1][j])

			// check if number continues to the left side of N
			k := 1
			for (j-k >= 0) && arr[i-1][j-k] != '.' {
				number = string(arr[i-1][j-k]) + number
				if k == 1 {
					coordinates["NW"] = true
				}
				k++
			}

			// check if number continues to the right side of N
			k = 1
			for (j+k < len(arr[0])) && arr[i-1][j+k] != '.' {
				number += string(arr[i-1][j+k])
				if k == 1 {
					coordinates["NE"] = true
				}
				k++
			}

			numberAsInt, _ := strconv.Atoi(number)
			numbers = append(numbers, numberAsInt)
		}

		if j != 0 {
			// check NW
			if unicode.IsDigit(arr[i-1][j-1]) && !coordinates["NW"] {
				coordinates["NW"] = true
				number := string(arr[i-1][j-1])

				// check if number continues to the left side of NW
				k := 1
				for (j-1-k >= 0) && arr[i-1][j-1-k] != '.' {
					number = string(arr[i-1][j-1-k]) + number
					k++
				}
				numberAsInt, _ := strconv.Atoi(number)
				numbers = append(numbers, numberAsInt)
				if len(numbers) > 2 {
					return 0
				}
			}
		}

		if j != len(arr[0])-1 {
			// check NE
			if unicode.IsDigit(arr[i-1][j+1]) && !coordinates["NE"] {
				coordinates["NE"] = true
				number := string(arr[i-1][j+1])

				// check if number continues to the right side of NE
				k := 1
				for (j+1+k < len(arr[0])) && arr[i-1][j+1+k] != '.' {
					number += string(arr[i-1][j+k+1])
					k++
				}
				numberAsInt, _ := strconv.Atoi(number)
				numbers = append(numbers, numberAsInt)
				if len(numbers) > 2 {
					return 0
				}
			}
		}
	}

	// check W
	if unicode.IsDigit(arr[i][j-1]) {
		coordinates["W"] = true
		number := string(arr[i][j-1])

		// check if number continues to the left side of W
		k := 1
		for (j-1-k >= 0) && arr[i][j-1-k] != '.' {
			number = string(arr[i][j-1-k]) + number
			k++
		}
		numberAsInt, _ := strconv.Atoi(number)
		numbers = append(numbers, numberAsInt)
		if len(numbers) > 2 {
			return 0
		}
	}

	// check E
	if unicode.IsDigit(arr[i][j+1]) {
		coordinates["E"] = true
		number := string(arr[i][j+1])

		// check if number continues to the right side of E
		k := 1
		for (j+1+k < len(arr[0])) && arr[i][j+1+k] != '.' {
			number += string(arr[i][j+1+k])
			k++
		}
		numberAsInt, _ := strconv.Atoi(number)
		numbers = append(numbers, numberAsInt)
		if len(numbers) > 2 {
			return 0
		}
	}

	if i != len(arr)-1 {
		// check S
		if unicode.IsDigit(arr[i+1][j]) {
			coordinates["S"] = true
			number := string(arr[i+1][j])

			// check if number continues to the left side of S
			k := 1
			for (j-k >= 0) && arr[i+1][j-k] != '.' {
				number = string(arr[i+1][j-k]) + number
				if k == 1 {
					coordinates["SW"] = true
				}
				k++
			}

			// check if number continues to the right side of S
			k = 1
			for (j+k < len(arr[0])) && arr[i+1][j+k] != '.' {
				number += string(arr[i+1][j+k])
				if k == 1 {
					coordinates["SE"] = true
				}
				k++
			}

			numberAsInt, _ := strconv.Atoi(number)
			numbers = append(numbers, numberAsInt)
			if len(numbers) > 2 {
				return 0
			}
		}

		if j != 0 {
			// check SW
			if unicode.IsDigit(arr[i+1][j-1]) && !coordinates["SW"] {
				coordinates["SW"] = true
				number := string(arr[i+1][j-1])

				// check if number continues to the left side of SW
				k := 1
				for (j-1-k >= 0) && arr[i+1][j-1-k] != '.' {
					number = string(arr[i+1][j-1-k]) + number
					k++
				}
				numberAsInt, _ := strconv.Atoi(number)
				numbers = append(numbers, numberAsInt)
				if len(numbers) > 2 {
					return 0
				}
			}
		}

		if j != len(arr[0])-1 {
			// check SE
			if unicode.IsDigit(arr[i+1][j+1]) && !coordinates["SE"] {
				coordinates["SE"] = true
				number := string(arr[i+1][j+1])

				// check if number continues to the right side of SE
				k := 1
				for (j+1+k < len(arr[0])) && arr[i+1][j+1+k] != '.' {
					number += string(arr[i+1][j+k+1])
					k++
				}
				numberAsInt, _ := strconv.Atoi(number)
				numbers = append(numbers, numberAsInt)
				if len(numbers) > 2 {
					return 0
				}
			}
		}
	}

	// for paranoia reasons another check
	if len(numbers) == 2 {
		fmt.Printf("%v\n", numbers)
		return numbers[0] * numbers[1]
	}

	return 0
}
