package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

type Module interface {
	ProcessSignal(int, string) int
	GetTargets() []string
}

type FlipFlop struct {
	State   bool
	Targets []string
}

func (ff *FlipFlop) ProcessSignal(pulse int, _ string) int {
	if pulse == 2 {
		return 0
	}

	if !ff.State {
		ff.State = true
		return 2
	} else {
		ff.State = false
		return 1
	}
}

func (ff FlipFlop) GetTargets() []string {
	return ff.Targets
}

type Conjunction struct {
	RecentPulses map[string]int
	Targets      []string
}

func (c *Conjunction) ProcessSignal(pulse int, sender string) int {
	c.RecentPulses[sender] = pulse
	for _, v := range c.RecentPulses {
		if v == 1 {
			return 2
		}
	}

	return 1
}

func (c Conjunction) GetTargets() []string {
	return c.Targets
}

type Broadcaster struct {
	Targets []string
}

func (b *Broadcaster) ProcessSignal(pulse int, _ string) int {
	return pulse
}

func (b Broadcaster) GetTargets() []string {
	return b.Targets
}

func main() {
	// int signal set to 2 means high pulse, 1 means low pulse, 0 means no pulse
	inputLines := utils.LoadPuzzleInput("../input.txt")

	modules := make(map[string]Module)
	for _, line := range inputLines {
		line = strings.ReplaceAll(line, " ", "")
		mod := strings.Split(line, "->")
		var module Module
		var name string
		if strings.HasPrefix(mod[0], "%") {
			module = &FlipFlop{
				State:   false,
				Targets: strings.Split(mod[1], ","),
			}
			name = mod[0][1:]
		} else if strings.HasPrefix(mod[0], "&") {
			module = &Conjunction{
				RecentPulses: make(map[string]int),
				Targets:      strings.Split(mod[1], ","),
			}
			name = mod[0][1:]
		} else {
			module = &Broadcaster{Targets: strings.Split(mod[1], ",")}
			name = mod[0]
		}
		modules[name] = module
	}
	// Initialize Conjunction modules memory to low pulses
	for k, v := range modules {
		for _, target := range v.GetTargets() {
			if mod, ok := modules[target].(*Conjunction); ok {
				mod.RecentPulses[k] = 1
			}
		}
	}

	highCounter := 0
	lowCounter := 0
	for range 1000 {
		fifo := make([]string, 0)

		fifo = append(fifo, "broadcaster,1,button")
		lowCounter++
	outer:
		for len(fifo) != 0 {
			curr := fifo[0]
			fifo = fifo[1:]
			modName := strings.Split(curr, ",")[0]
			signal, _ := strconv.Atoi(strings.Split(curr, ",")[1])
			sender := strings.Split(curr, ",")[2]

			signal = modules[modName].ProcessSignal(signal, sender)
			for range len(modules[modName].GetTargets()) {
				switch signal {
				case 0:
					continue outer
				case 1:
					lowCounter++
				case 2:
					highCounter++
				}
			}
			for _, target := range modules[modName].GetTargets() {
				found := false
				for k := range modules {
					if k == target {
						found = true
					}
				}
				if !found {
					continue
				}
				fifo = append(fifo, target+","+strconv.Itoa(signal)+","+modName)
			}
		}
	}
	fmt.Printf("Pulses: %d", highCounter*lowCounter)
}
