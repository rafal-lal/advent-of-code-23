package main

import (
	// "fmt"
	// "strings"

	"fmt"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	pattern := make([][]string, 0)
	startPoint := [2]int{}
	for i, line := range inputLines {
		pattern = append(pattern, []string{})
		for j, tile := range line {
			pattern[i] = append(pattern[i], string(tile))
			if string(tile) == "S" {
				startPoint[0] = i
				startPoint[1] = j
			}
		}
	}

	current := "|"
	source := "N"
	steps := 0
	point := [2]int{startPoint[0] + 1, startPoint[1]}
	for {
		direction := leadsTo(current, source)

		switch direction {
		case "N":
			point[0]--
			source = "S"
		case "E":
			point[1]++
			source = "W"
		case "S":
			point[0]++
			source = "N"
		case "W":
			point[1]--
			source = "E"
		}

		current = pattern[point[0]][point[1]]
		steps++

		if point == startPoint {
			break
		}
	}

	fmt.Printf("Steps to farthest point %d\n", (steps+1)/2)

}

func leadsTo(pipe, source string) string {
	switch pipe {
	case "|":
		if source == "N" {
			return "S"
		} else if source == "S" {
			return "N"
		}
	case "-":
		if source == "E" {
			return "W"
		} else if source == "W" {
			return "E"
		}
	case "L":
		if source == "N" {
			return "E"
		} else if source == "E" {
			return "N"
		}
	case "J":
		if source == "N" {
			return "W"
		} else if source == "W" {
			return "N"
		}
	case "7":
		if source == "S" {
			return "W"
		} else if source == "W" {
			return "S"
		}
	case "F":
		if source == "S" {
			return "E"
		} else if source == "E" {
			return "S"
		}
	default:
		return ""
	}

	return ""
}
