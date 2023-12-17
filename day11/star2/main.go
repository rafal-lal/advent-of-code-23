package main

import (
	"container/list"
	"fmt"

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
	replace := make([]string, 0)
	for i := 0; i < len(universe[0]); i++ {
		replace = append(replace, "*")
	}

	for _, idx := range rowsToEnlarge {
		universe[idx] = replace
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

	for _, idx := range colsToEnlarge {
		for i := range universe {
			universe[i][idx] = "*"
		}

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

	var sum int64 = 0
    occ := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
            if occ % 1000 == 0 {
                println(occ)
            }
			sum += shortestPath(universe, galaxies[i], galaxies[j])
            occ++
		}
	}

	fmt.Printf("Sum of distances %d\n", sum)
}

// Weighted BFS algorithm, non optimal, could have used Dijkstra's
func shortestPath(grid [][]string, A, B [2]int) int64 {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := list.New()
	queue.PushBack(struct {
		point    [2]int
		distance int64
	}{A, 0})
	visited[A[0]][A[1]] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(struct {
			point    [2]int
			distance int64
		})

		if current.point == B {
			return current.distance
		}

		neighbors := getNeighbors(current.point[0], current.point[1], rows, cols)
		for _, neighbor := range neighbors {
			neighborI, neighborJ := neighbor[0], neighbor[1]
			if !visited[neighborI][neighborJ] {
				weight := 1
				if grid[neighborI][neighborJ] == "*" {
					weight = 1000000
				}
				queue.PushBack(struct {
					point    [2]int
					distance int64
				}{neighbor, current.distance + int64(weight)})
				visited[neighborI][neighborJ] = true
			}
		}
	}

	return -1
}

func getNeighbors(i, j, rows, cols int) [][2]int {
	neighbors := [][2]int{}
	if i > 0 {
		neighbors = append(neighbors, [2]int{i - 1, j})
	}
	if i < rows-1 {
		neighbors = append(neighbors, [2]int{i + 1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, [2]int{i, j - 1})
	}
	if j < cols-1 {
		neighbors = append(neighbors, [2]int{i, j + 1})
	}

	return neighbors
}
