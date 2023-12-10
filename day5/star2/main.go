package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

// just brute forcing with concurrency
func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	seedsAsStr := strings.Fields(inputLines[0])
	seedsAsStr = seedsAsStr[1:]
	lowestLocation := 2147483647 // highest int

	for i := 0; i < len(seedsAsStr); i += 2 {
		var seeds []int
		data := make([][][]int, 0)

		val, _ := strconv.Atoi(seedsAsStr[i])
		rang, _ := strconv.Atoi(seedsAsStr[i+1])
		for j := val; j < val+rang; j++ {
			seeds = append(seeds, j)
		}
		

		sectionCounter := 0
		lineCounter := 0
		data = append(data, [][]int{})
		for _, line := range inputLines[3:] {
			if line == "" {
				continue
			}
			lineHasNum := regexp.MustCompile(`\d`).MatchString(line)
			if !lineHasNum {
				sectionCounter++
				data = append(data, [][]int{})
				lineCounter = 0
				continue
			}

			data[sectionCounter] = append(data[sectionCounter], []int{})
			numsAsStr := strings.Fields(line)
			for _, numAsStr := range numsAsStr {
				num, _ := strconv.Atoi(numAsStr)
				data[sectionCounter][lineCounter] = append(data[sectionCounter][lineCounter], num)
			}
			lineCounter++
		}

		calculatedLocations := make([]int, len(seeds))
		
		initBatch := int(math.Floor(float64(len(seeds)) / 8.0))
		mod := len(seeds) % 8
		
		var wg sync.WaitGroup

		for i := 0; i < 8; i++ {
			batch := initBatch
			if i == 7 {
				batch += mod
			}

			wg.Add(1)
			go func(i, initBatch, batch int) {
				defer wg.Done()
				for j := i * initBatch; j < i*initBatch+batch; j++ {
					lastVal := seeds[j]
					for _, outerVal := range data {
						currentVal := -1
						for _, val := range outerVal {
							if lastVal >= val[1] && lastVal < (val[1]+val[2]) {
								currentVal = lastVal + (val[0] - val[1])
							}
						}
						if currentVal == -1 {
							currentVal = lastVal
						}
						lastVal = currentVal
					}

					location := lastVal
					calculatedLocations[j] = location
				}
			}(i, initBatch, batch)
		}
		wg.Wait()

		lowestLocationLocal := 2147483647 // highest int
		for _, val := range calculatedLocations {
			if val < lowestLocationLocal {
				lowestLocationLocal = val
			}
		}
		println(lowestLocationLocal)
		if lowestLocationLocal < lowestLocation {
			lowestLocation = lowestLocationLocal
		}
	}

	fmt.Printf("Lowest location: %d\n", lowestLocation)
}
