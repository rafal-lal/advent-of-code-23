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

	currentKey := "AAA"
	counter := 0

	i := 0
	for {
		counter++
		if string(sequence[i]) == "R" {
			currentKey = nodes[currentKey][1]
		} else if string(sequence[i]) == "L" {
			currentKey = nodes[currentKey][0]
		}

		if currentKey == "ZZZ" {
			break
		}

		i++

		if i == len(sequence) {
			i = 0
		}
	}

	fmt.Printf("Required steps: %d\n", counter)
}
