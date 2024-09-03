package main

import (
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	var validCombs int
	for _, line := range inputLines {
		chars := strings.Split(strings.Split(line, " ")[0], "")
		nums := strings.Split(strings.Split(line, " ")[1], ",")

		numOfQuest := 0
		for _, char := range chars {
			if char == "?" {
				numOfQuest++
			}
		}

		combs := generateCombinations([]string{"#", "."}, numOfQuest)
		var el string
		for _, comb := range combs {
			combSplit := strings.Split(comb, "")

			charsCopy := make([]string, len(chars))
			copy(charsCopy, chars)

			for i, char := range charsCopy {
				if char == "?" {
					el, combSplit = combSplit[0], combSplit[1:]
					charsCopy[i] = el
				}
			}

			if check(charsCopy, nums) {
				validCombs++
			}
		}
	}

	println(validCombs)
}

func check(comb, nums []string) bool {
	numsAsInt := make([]int, len(nums))
	for i, num := range nums {
		numsAsInt[i], _ = strconv.Atoi(num)
	}

	hashGroupLength := 0
	hashGroups := make([]int, 0)
	for _, ch := range comb {
		if ch == "#" {
			hashGroupLength++
			continue
		}
		if hashGroupLength != 0 {
			hashGroups = append(hashGroups, hashGroupLength)
		}
		hashGroupLength = 0
	}
	if hashGroupLength != 0 {
		hashGroups = append(hashGroups, hashGroupLength)
	}

	if len(numsAsInt) != len(hashGroups) {
		return false
	}

	for i := range hashGroups {
		if hashGroups[i] != numsAsInt[i] {
			return false
		}
	}

	return true
}

func generateCombinations(chars []string, places int) []string {
	if places == 0 {
		return []string{""}
	}

	var combinations []string
	subCombinations := generateCombinations(chars, places-1)

	for _, char := range chars {
		for _, subComb := range subCombinations {
			combinations = append(combinations, char+subComb)
		}
	}

	return combinations
}
