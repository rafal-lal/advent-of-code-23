package main

import (
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

type Pair struct {
	Label       string
	FocalLength int
}

func Remove(boxes []Pair, label string) []Pair {
	for i, box := range boxes {
		if box.Label == label {
			return append(boxes[:i], boxes[i+1:]...)
		}
	}

	return boxes
}

func Add(boxes []Pair, p *Pair) []Pair {
	found := false
	for i, box := range boxes {
		if box.Label == p.Label {
			boxes[i] = *p
			found = true
		}
	}

	if !found {
		boxes = append(boxes, *p)
	}

	return boxes
}

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")
	steps := strings.Split(inputLines[0], ",")

	boxes := make([][]Pair, 256)
	for _, step := range steps {
		hash := 0
		focal := -1
		label := ""
		for i, ch := range step {
			if string(ch) == "=" {
				focal, _ = strconv.Atoi(string(step[i+1]))
				break
			} else if string(ch) == "-" {
				break
			}
			hash += int(ch)
			hash *= 17
			hash %= 256
			label += string(ch)
		}

		if focal != -1 {
			boxes[hash] = Add(boxes[hash], &Pair{label, focal})
		} else {
			boxes[hash] = Remove(boxes[hash], label)
		}

	}

	sum := 0
	for i, box := range boxes {
		for j, pair := range box {
			sum += (1 + i) * (j+1) * pair.FocalLength
		}
	}

	println(sum)
}
