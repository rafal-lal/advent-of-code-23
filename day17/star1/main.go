package main

import (
	"reflect"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

}

type Point struct {
	Y int
	X int
}

func Dijkstra(grid [][]int, source Point) {
	dist := make(map[Point]int)
	dist[source] = 0

	queue := make([]Point, 0)

	for i := range grid {
		for j := range grid[0] {
			point := Point{i, j}
			if !reflect.DeepEqual(source, point) {
				dist[point] = 2147483647
			}
			queue = append(queue, point)
		}
	}

	for len(queue) > 0 {
		curr := findMinPoint(dist)
		for i, val := range queue {
			if reflect.DeepEqual(val, curr) {
				queue = utils.RemoveOrdered(queue, i)
				break
			}
		}



	}

}

func findMinPoint(dist map[Point]int) Point {
	min := 2147483647
	var minPoint Point
	for k, v := range dist {
		if v < min {
			min = v
			minPoint = k
		}
	}

	return minPoint
}

func findNeighbor() {
	// neighbors have limitations - can turn left, right or continue straigt but not for more than 3 blocks
}