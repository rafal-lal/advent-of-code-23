package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

type Rule struct {
	Letter string
	Sign   string
	Number int
	Target string
}

type Part map[string]int

func (r *Rule) ProcessPart(p Part) (bool, string) {
	if r.Number == -1 {
		return false, r.Target
	}

	var next bool = true
	switch r.Sign {
	case ">":
		if p[r.Letter] > r.Number {
			next = false
		}
	case "<":
		if p[r.Letter] < r.Number {
			next = false
		}
	case "=":
		if p[r.Letter] == r.Number {
			next = false
		}
	}

	return next, r.Target
}

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	workflow := make(map[string][]Rule, 100)
	var idx int
	for i, line := range inputLines {
		if line == "" {
			idx = i
			break
		}
		name := strings.Split(line, "{")[0]
		rest := strings.Split(line, "{")[1]
		rest = rest[:len(rest)-1]
		rules := strings.Split(rest, ",")
		rulesSlice := make([]Rule, 0)
		for _, rule := range rules {
			if !strings.ContainsAny(rule, "<>=") {
				rulesSlice = append(rulesSlice, Rule{"", "", -1, rule})
				continue
			}
			numberStr := strings.Split(rule, ":")[0][2:]
			number, _ := strconv.Atoi(numberStr)
			ruleStruct := Rule{
				Letter: string(rule[0]),
				Sign:   string(rule[1]),
				Number: number,
				Target: strings.Split(rule, ":")[1],
			}
			rulesSlice = append(rulesSlice, ruleStruct)
		}
		workflow[name] = rulesSlice
	}

	parts := make([]Part, 0, 100)
	for _, line := range inputLines[idx+1:] {
		nums, _ := strings.CutPrefix(line, "{")
		nums, _ = strings.CutSuffix(nums, "}")
		numsSlice := strings.Split(nums, ",")
		x, _ := strconv.Atoi(numsSlice[0][2:])
		m, _ := strconv.Atoi(numsSlice[1][2:])
		a, _ := strconv.Atoi(numsSlice[2][2:])
		s, _ := strconv.Atoi(numsSlice[3][2:])
		parts = append(parts, Part{"x": x, "m": m, "a": a, "s": s})
	}

	sum := 0
	for _, part := range parts {
		ruleId := 0
		wf := "in"
		wfToBe := ""
		next := false
		for wf != "A" && wf != "R" {
			next, wfToBe = workflow[wf][ruleId].ProcessPart(part)
			if next {
				ruleId++
				continue
			}
			wf = wfToBe
			ruleId = 0
		}
		if wf == "A" {
			sum += part["x"] + part["m"] + part["a"] + part["s"]
		}
	}
	fmt.Printf("Sum : %d", sum)
}
