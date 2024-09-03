package main

import (
	"reflect"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	patterns := make([][][]string, 0)
	patterns = append(patterns, [][]string{})
	cnt := 0
	cntInternal := 0
	for _, line := range inputLines {
		if line == "" {
			patterns = append(patterns, [][]string{})
			cntInternal = 0
			cnt++
			continue
		}

		patterns[cnt] = append(patterns[cnt], []string{})
		for _, char := range line {
			patterns[cnt][cntInternal] = append(patterns[cnt][cntInternal], string(char))
		}
		cntInternal++
	}

	sum := 0
	for _, pattern := range patterns {
		if hor := findHorizontal(pattern); hor != -1 {
			sum += 100 * (hor + 1)
		}
		if ver := findVertical(pattern); ver != -1 {
			sum += ver + 1
		}
	}

	println(sum)
}

func findHorizontal(pattern [][]string) int {
	validLines := make([]int, 0)
	for i := 0; i < len(pattern)-1; i++ {
		if reflect.DeepEqual(pattern[i], pattern[i+1]) {
			validLines = append(validLines, i)
		}
	}

	if len(validLines) == 0 {
		return -1
	}

	for _, validLine := range validLines {
		valid := true

		up := validLine + 2
		down := validLine - 1
		for up < len(pattern) && down >= 0 {
			if !reflect.DeepEqual(pattern[up], pattern[down]) {
				valid = false
				break
			}
			up++
			down--
		}
		if valid {
			return validLine
		}

	}

	return -1
}

func findVertical(pattern [][]string) int {
	validLines := make([]int, 0)
	for i := 0; i < len(pattern[0])-1; i++ {
		same := true
		for j := 0; j < len(pattern); j++ {
			if pattern[j][i] != pattern[j][i+1] {
				same = false
			}
		}
		if same {
			validLines = append(validLines, i)
		}
	}

	if len(validLines) == 0 {
		return -1
	}

	for _, validLine := range validLines {
		valid := true

		left := validLine - 1
		right := validLine + 2

		for left >= 0 && right < len(pattern[0]) {
			same := true
			for i := 0; i < len(pattern); i++ {
				if pattern[i][left] != pattern[i][right] {
					same = false
				}
			}
			if !same {
				valid = false
				break
			}
			left--
			right++
		}
		if valid {
			return validLine
		}
	}

	return -1
}
