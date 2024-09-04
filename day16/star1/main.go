package main

import (
	// "fmt"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

type BeamHead struct {
	X         int
	Y         int
	Direction string
}

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	grid := make([][]string, 0)
	for i, line := range inputLines {
		grid = append(grid, []string{})
		for _, ch := range line {
			grid[i] = append(grid[i], string(ch))
		}
	}

	energizedMap := make([][]bool, 0)
	for range len(grid) {
		energizedRow := make([]bool, len(grid[0]))
		energizedMap = append(energizedMap, energizedRow)
	}

	beams := []*BeamHead{{0, 0, "right"}}
	energizedMap[beams[0].Y][beams[0].X] = true
	cntt := 0
	for len(beams) > 0 && cntt < 100000 {
		cntt++
		cnt := 0
		for i := range energizedMap {
			for j := range energizedMap[0] {
				if energizedMap[i][j] {
					cnt++
				}
			}
		}
		println(cnt)

		beamsIdxToDel := make([]int, 0)
		beamsToAdd := make([]*BeamHead, 0)

		for i, beam := range beams {
			if bh := moveBeam(grid, beam); bh != nil {
				if !isOutOfGrid(grid, bh) {
					beamsToAdd = append(beamsToAdd, bh)
				}
			}
			if isOutOfGrid(grid, beam) {
				beamsIdxToDel = append(beamsIdxToDel, i)
			} else {
				energizedMap[beam.Y][beam.X] = true
			}
		}

		for _, idx := range beamsIdxToDel {
			beams = utils.RemoveOrdered(beams, idx)
			for j, i := range beamsIdxToDel {
				beamsIdxToDel[j] = i - 1
			}
		}

		for _, beamToAdd := range beamsToAdd {
			energizedMap[beamToAdd.Y][beamToAdd.X] = true
		}
		beams = append(beams, beamsToAdd...)
	}

	cnt := 0
	for i := range energizedMap {
		for j := range energizedMap[0] {
			if energizedMap[i][j] {
				cnt++
			}
		}
	}
	println(cnt)
}

func moveBeam(grid [][]string, c *BeamHead) *BeamHead {
	switch c.Direction {
	case "up":
		switch grid[c.Y][c.X] {
		case ".":
			c.Y -= 1
			return nil
		case "|":
			c.Y -= 1
			return nil
		case "-":
			c.X -= 1
			c.Direction = "left"
			return &BeamHead{c.X + 2, c.Y, "right"}
		case "\\":
			c.X -= 1
			c.Direction = "left"
			return nil
		case "/":
			c.X += 1
			c.Direction = "right"
			return nil
		}

	case "right":
		switch grid[c.Y][c.X] {
		case ".":
			c.X += 1
			return nil
		case "|":
			c.Y -= 1
			c.Direction = "up"
			return &BeamHead{c.X, c.Y + 2, "down"}
		case "-":
			c.X += 1
			return nil
		case "\\":
			c.Y += 1
			c.Direction = "down"
			return nil
		case "/":
			c.Y -= 1
			c.Direction = "up"
			return nil
		}

	case "down":
		switch grid[c.Y][c.X] {
		case ".":
			c.Y += 1
			return nil
		case "|":
			c.Y += 1
			return nil
		case "-":
			c.X -= 1
			c.Direction = "left"
			return &BeamHead{c.X + 2, c.Y, "right"}
		case "\\":
			c.X += 1
			c.Direction = "right"
			return nil
		case "/":
			c.X -= 1
			c.Direction = "left"
			return nil
		}

	case "left":
		switch grid[c.Y][c.X] {
		case ".":
			c.X -= 1
			return nil
		case "|":
			c.Y -= 1
			c.Direction = "up"
			return &BeamHead{c.X, c.Y + 2, "down"}
		case "-":
			c.X -= 1
			return nil
		case "\\":
			c.Y -= 1
			c.Direction = "up"
			return nil
		case "/":
			c.Y += 1
			c.Direction = "down"
			return nil
		}
	}

	return nil
}

func isOutOfGrid(grid [][]string, b *BeamHead) bool {
	return b.Y < 0 || b.Y > len(grid)-1 || b.X < 0 || b.X > len(grid[0])-1
}
