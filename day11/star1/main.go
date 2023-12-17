package main

import (
	// "fmt"
	"fmt"
	"math"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	universe := make([][]string, 0)
	for i, line := range inputLines {
		universe = append(universe, []string{})
		for _, dot := range line {
			universe[i] = append(universe[i], string(dot))
		}
	}

	// mark rows that need enlarging
	rowsToEnlarge := make([]int, 0)
	for i, row := range universe {
		emptyRow := true
		for _, char := range row {
			if char == "#" {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			rowsToEnlarge = append(rowsToEnlarge, i)
		}
	}
	toAppend := make([]string, 0)
	for i := 0; i < len(universe[0]); i++ {
		toAppend = append(toAppend, ".")
	}

	counter := 0
	for _, idx := range rowsToEnlarge {
		idx += counter
		universe = append(universe[:idx+1], universe[idx:]...)
		universe[idx] = toAppend
		counter++
	}

	// mark columns that need enlarging
	colsToEnlarge := make([]int, 0)
	for j := 0; j < len(universe[0]); j++ {
		emptyCol := true
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == "#" {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			colsToEnlarge = append(colsToEnlarge, j)
		}
	}

	counter = 0
	for _, idx := range colsToEnlarge {
		idx += counter
		for i := range universe {
			universe[i] = append(universe[i][:idx+1], universe[i][idx:]...)
			universe[i][idx] = "."
		}
		counter++
	}

	// mark galaxies locations
	galaxies := make([][2]int, 0)
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == "#" {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	sum := 0
	occ := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += manhattanDistance(galaxies[i], galaxies[j])
			occ++
		}
	}

	fmt.Printf("Sum of distances %d\n", sum)
}

func manhattanDistance(A, B [2]int) int {
	return int(math.Abs(float64(A[0]-B[0])) + math.Abs(float64(A[1]-B[1])))
}
