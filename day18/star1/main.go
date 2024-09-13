package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

type Step struct {
	Direction string
	Number    int
	Color     string
}

type Position struct {
	Row int
	Col int
}

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	steps := make([]Step, 0)
	for _, line := range inputLines {
		direction := strings.Split(line, " ")[0]
		numberStr := strings.Split(line, " ")[1]
		number, _ := strconv.Atoi(numberStr)
		color := strings.Split(line, " ")[2]
		color = strings.TrimPrefix(color, "(")
		color = strings.TrimSuffix(color, ")")

		steps = append(steps, Step{direction, number, color})
	}

	pos := Position{0, 0}
	grid := make([][]string, 0)
	grid = append(grid, []string{"+"})
	for _, step := range steps {
		switch step.Direction {
		case "U":
			pos.Row -= 1
			if pos.Row-step.Number < 0 {
				rowsToAdd := -(pos.Row - step.Number + 1)
				pos.Row += rowsToAdd
				for range rowsToAdd {
					emptyRow := make([]string, len(grid[0]))
					grid = slices.Insert(grid, 0, emptyRow)
				}
			}
			for i := pos.Row; i > pos.Row-step.Number; i-- {
				grid[i][pos.Col] = step.Color
			}
			pos.Row -= step.Number
			pos.Row += 1
		case "R":
			pos.Col += 1
			if pos.Col+step.Number >= len(grid[0]) {
				for i := range grid {
					for range (pos.Col + step.Number) - len(grid[i]) {
						grid[i] = append(grid[i], "")
					}
				}
			}
			for i := pos.Col; i < pos.Col+step.Number; i++ {
				grid[pos.Row][i] = step.Color
			}
			pos.Col += step.Number - 1
		case "D":
			pos.Row += 1
			if pos.Row+step.Number >= len(grid) {
				for range (pos.Row + step.Number) - len(grid) {
					emptyRow := make([]string, len(grid[0]))
					grid = append(grid, emptyRow)
				}
			}
			for i := pos.Row; i < pos.Row+step.Number; i++ {
				grid[i][pos.Col] = step.Color
			}
			pos.Row += step.Number - 1
		case "L":
			pos.Col -= 1
			if pos.Col-step.Number < 0 {
				colsToAdd := -(pos.Col - step.Number + 1)
				pos.Col += colsToAdd
				for i := range grid {
					for range colsToAdd {
						grid[i] = slices.Insert(grid[i], 0, "")
					}
				}
			}
			for i := pos.Col; i > pos.Col-step.Number; i-- {
				grid[pos.Row][i] = step.Color
			}
			pos.Col -= step.Number
			pos.Col += 1
		}
	}
	printGrid(grid)
	floodGrid(grid)
	printGrid(grid)

	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "" {
				cnt++	
			}
		}
	}
	println(cnt)
}

func floodGrid(grid [][]string) {
	floodCol := func(dir string, pos Position) {
		for grid[pos.Row][pos.Col] == "" || grid[pos.Row][pos.Col] == "+" {
			grid[pos.Row][pos.Col] = "+"
			if dir == "U" {
				pos.Row--
			} else {
				pos.Row++
			}
		}
	}
	floodRow := func(dir string, pos Position) {
		for grid[pos.Row][pos.Col] == "" || grid[pos.Row][pos.Col] == "+" {
			grid[pos.Row][pos.Col] = "+"
			if dir == "L" {
				pos.Col--
			} else {
				pos.Col++
			}
		}
	}
	checkNextMove := func(currDir string, pos Position) (Position, string) {
		switch currDir {
		case "R":
			// continue right
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col + 1}) && strings.HasPrefix(grid[pos.Row][pos.Col+1], "#") {
				return Position{pos.Row, pos.Col + 1}, "R"
			}
			// switch up
			if !isOutOfGrid(grid, &Position{pos.Row - 1, pos.Col}) && strings.HasPrefix(grid[pos.Row-1][pos.Col], "#") {
				return Position{pos.Row - 1, pos.Col}, "U"
			}
			// switch down
			if !isOutOfGrid(grid, &Position{pos.Row + 1, pos.Col}) && strings.HasPrefix(grid[pos.Row+1][pos.Col], "#") {
				return Position{pos.Row + 1, pos.Col}, "D"
			}
		case "D":
			// continue down
			if !isOutOfGrid(grid, &Position{pos.Row + 1, pos.Col}) && strings.HasPrefix(grid[pos.Row+1][pos.Col], "#") {
				return Position{pos.Row + 1, pos.Col}, "D"
			}
			// switch left
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col - 1}) && strings.HasPrefix(grid[pos.Row][pos.Col-1], "#") {
				return Position{pos.Row, pos.Col - 1}, "L"
			}
			// switch right
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col + 1}) && strings.HasPrefix(grid[pos.Row][pos.Col+1], "#") {
				return Position{pos.Row, pos.Col + 1}, "R"
			}
		case "L":
			// continue left
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col - 1}) && strings.HasPrefix(grid[pos.Row][pos.Col-1], "#") {
				return Position{pos.Row, pos.Col - 1}, "L"
			}
			// switch up
			if !isOutOfGrid(grid, &Position{pos.Row - 1, pos.Col}) && strings.HasPrefix(grid[pos.Row-1][pos.Col], "#") {
				return Position{pos.Row - 1, pos.Col}, "U"
			}
			// switch down
			if !isOutOfGrid(grid, &Position{pos.Row + 1, pos.Col}) && strings.HasPrefix(grid[pos.Row+1][pos.Col], "#") {
				return Position{pos.Row + 1, pos.Col}, "D"
			}
		case "U":
			// continue up
			if !isOutOfGrid(grid, &Position{pos.Row - 1, pos.Col}) && strings.HasPrefix(grid[pos.Row-1][pos.Col], "#") {
				return Position{pos.Row - 1, pos.Col}, "U"
			}
			// switch left
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col - 1}) && strings.HasPrefix(grid[pos.Row][pos.Col-1], "#") {
				return Position{pos.Row, pos.Col - 1}, "L"
			}
			// switch right
			if !isOutOfGrid(grid, &Position{pos.Row, pos.Col + 1}) && strings.HasPrefix(grid[pos.Row][pos.Col+1], "#") {
				return Position{pos.Row, pos.Col + 1}, "R"
			}
		}

		return Position{}, ""
	}

	var startPoint Position
	var currPoint Position
	direction := "R"
	for i := range grid[0] {
		if grid[0][i] != "" {
			startPoint = Position{0, i}
			break
		}
	}
	currPoint = Position{startPoint.Row, startPoint.Col + 1}
	cnt := 0
	for currPoint.Row != startPoint.Row || currPoint.Col != startPoint.Row {
		switch direction {
		case "R":
			floodCol("D", Position{currPoint.Row + 1, currPoint.Col})
		case "D":
			floodRow("L", Position{currPoint.Row, currPoint.Col - 1})
		case "L":
			floodCol("U", Position{currPoint.Row - 1, currPoint.Col})
		case "U":
			floodRow("R", Position{currPoint.Row, currPoint.Col + 1})
		}
		currPoint, direction = checkNextMove(direction, currPoint)
		// printGrid(grid)
		cnt++
		if cnt == 10000 {
			break
		}
	}
}

func printGrid(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "+" {
				fmt.Printf("+")
			} else if grid[i][j] != "" {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func isOutOfGrid(grid [][]string, p *Position) bool {
	return p.Row < 0 || p.Row > len(grid)-1 || p.Col < 0 || p.Col > len(grid[0])-1
}
