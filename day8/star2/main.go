package main

import (
	"fmt"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	sequence := inputLines[0]
	nodes := make(map[string][2]string)

	for _, line := range inputLines[2:] {
		line = strings.Replace(line, "=", "", -1)
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)
		line = strings.Replace(line, ",", "", -1)

		fields := strings.Fields(line)
		nodes[fields[0]] = [2]string{fields[1], fields[2]}
	}

	currentKeys := make([]string, 0)
	for key := range nodes {
		if strings.HasSuffix(key, "A") {
			currentKeys = append(currentKeys, key)
		}
	}

	stoppingPoints := make([]int, 0)
	counter := 0

	i := 0
	for {
		counter++
		for j := range currentKeys {
			if string(sequence[i]) == "R" {
				currentKeys[j] = nodes[currentKeys[j]][1]
			} else if string(sequence[i]) == "L" {
				currentKeys[j] = nodes[currentKeys[j]][0]
			}
		}

		for _, val := range currentKeys {
			if strings.HasSuffix(val, "Z") {
				stoppingPoints = append(stoppingPoints, counter)
			}
		}

		if len(stoppingPoints) == len(currentKeys) {
			break
		}

		i++

		if i == len(sequence) {
			i = 0
		}
	}

	fmt.Printf("Required steps: %d\n", LCM(stoppingPoints[0], stoppingPoints[1], stoppingPoints[2:]...))
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
