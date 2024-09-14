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

type Range struct {
	Vals     map[string][]int
	Workflow string
	Idx      int
}

func (r *Range) DeepCopy() *Range {
	dest := new(Range)
	dest.Vals = make(map[string][]int)
	for k, v := range r.Vals {
		slc := make([]int, len(v))
		copy(slc, v)
		dest.Vals[k] = slc
	}
	dest.Idx = r.Idx
	dest.Workflow = r.Workflow

	return dest
}

func (r *Rule) ProcessPart(rang Range) (Range, Range, string, bool) {
	if r.Number == -1 {
		return rang, rang, r.Target, true
	}

	passed := make([]int, 0)
	notPassed := make([]int, 0)
	switch r.Sign {
	case ">":
		for i := rang.Vals[r.Letter][0]; i < rang.Vals[r.Letter][len(rang.Vals[r.Letter])-1]+1; i++ {
			if i > r.Number {
				passed = append(passed, i)
				continue
			}
			notPassed = append(notPassed, i)
		}
	case "<":
		for i := rang.Vals[r.Letter][0]; i < rang.Vals[r.Letter][len(rang.Vals[r.Letter])-1]+1; i++ {
			if i < r.Number {
				passed = append(passed, i)
				continue
			}
			notPassed = append(notPassed, i)
		}
	case "=":
		for i := rang.Vals[r.Letter][0]; i < rang.Vals[r.Letter][len(rang.Vals[r.Letter])-1]+1; i++ {
			if i == r.Number {
				passed = append(passed, i)
				continue
			}
			notPassed = append(notPassed, i)
		}
	}
	notPassedRange := rang.DeepCopy()
	passedRange := rang.DeepCopy()
	notPassedRange.Vals[r.Letter] = notPassed
	passedRange.Vals[r.Letter] = passed

	return *notPassedRange, *passedRange, r.Target, false
}

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	workflow := make(map[string][]Rule, 100)
	for _, line := range inputLines {
		if line == "" {
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

	x := make([]int, 4000)
	m := make([]int, 4000)
	a := make([]int, 4000)
	s := make([]int, 4000)
	for i := range 4000 {
		x[i] = i + 1
		m[i] = i + 1
		a[i] = i + 1
		s[i] = i + 1
	}
	rang := Range{map[string][]int{"x": x, "m": m, "a": a, "s": s}, "in", 0}
	queue := make([]Range, 1)
	queue[0] = rang

	sum := 0
	numOfAppends := 0
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		notPassed, passed, target, both := workflow[curr.Workflow][curr.Idx].ProcessPart(curr)
		// if both is true there was no split
		if both {
			if target == "A" {
				sum += len(passed.Vals["x"]) * len(passed.Vals["m"]) * len(passed.Vals["a"]) * len(passed.Vals["s"])
				fmt.Printf("%d,%d %d,%d %d,%d %d,%d\n",
					passed.Vals["x"][0],
					passed.Vals["x"][len(passed.Vals["x"])-1],
					passed.Vals["m"][0], passed.Vals["m"][len(passed.Vals["m"])-1],
					passed.Vals["a"][0], passed.Vals["a"][len(passed.Vals["a"])-1],
					passed.Vals["s"][0], passed.Vals["s"][len(passed.Vals["s"])-1])
				numOfAppends++
				continue
			}
			if target == "R" {
				continue
			}
			passed.Workflow = target
			passed.Idx = 0
			queue = append(queue, passed)
			continue
		}
		if target == "A" {
			sum += len(passed.Vals["x"]) * len(passed.Vals["m"]) * len(passed.Vals["a"]) * len(passed.Vals["s"])
			fmt.Printf("%d,%d %d,%d %d,%d %d,%d\n",
				passed.Vals["x"][0],
				passed.Vals["x"][len(passed.Vals["x"])-1],
				passed.Vals["m"][0], passed.Vals["m"][len(passed.Vals["m"])-1],
				passed.Vals["a"][0], passed.Vals["a"][len(passed.Vals["a"])-1],
				passed.Vals["s"][0], passed.Vals["s"][len(passed.Vals["s"])-1])
			numOfAppends++
			notPassed.Idx++
			queue = append(queue, notPassed)
			continue
		}
		if target == "R" {
			notPassed.Idx++
			queue = append(queue, notPassed)
			continue
		}
		passed.Workflow = target
		passed.Idx = 0
		queue = append(queue, passed)
		notPassed.Idx++
		queue = append(queue, notPassed)
	}

	fmt.Printf("Sum : %d\n", sum)
	fmt.Printf("Sum : %d\n", numOfAppends)
}
