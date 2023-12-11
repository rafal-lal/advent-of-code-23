package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rafal-lal/advent-of-code-23/utils"
)

func main() {
	inputLines := utils.LoadPuzzleInput("../input.txt")

	ratedHands := make(map[string][]string, 7)

	// rate each hand and put them into seperate slices
	for _, line := range inputLines {
		ratedHands[rateHand(line)] = append(ratedHands[rateHand(line)], line)
	}

	// give value to every card
	cardsVals := map[string]int{
		"A": 13, "K": 12, "Q": 11,
		"J": 10, "T": 9, "9": 8,
		"8": 7, "7": 6, "6": 5, "5": 4,
		"4": 3, "3": 2, "2": 1,
	}

	// sort every rating from lowest to highest
	for _, rating := range ratedHands {
		sort.SliceStable(rating, func(i, j int) bool {
			handI := strings.Fields(rating[i])[0]
			handJ := strings.Fields(rating[j])[0]
			for k := 0; k < 5; k++ {
				if cardsVals[string(handI[k])] > cardsVals[string(handJ[k])] {
					return false
				}
				if cardsVals[string(handI[k])] < cardsVals[string(handJ[k])] {
					return true
				}
			}
			// insignificant
			return true
		})
	}

	// calculate total winning
	totalWinnings := 0
	rank := 1
	for _, key := range []string{"highCard", "onePair", "twoPair", "threeKind",
		"fullHouse", "fourKind", "fiveKind"} {
		for _, hand := range ratedHands[key] {
			valOfHandAsStr := strings.Fields(hand)[1]
			valOfHand, _ := strconv.Atoi(valOfHandAsStr)
			valOfHand *= rank
			totalWinnings += valOfHand
			rank++
		}
	}

	fmt.Printf("Total winnings %d\n", totalWinnings)
}

func rateHand(line string) string {
	hand := strings.Fields(line)[0]

	labels := map[string]int{
		string(hand[0]): 0,
		string(hand[1]): 0,
		string(hand[2]): 0,
		string(hand[3]): 0,
		string(hand[4]): 0,
	}

	for _, val := range hand {
		labels[string(val)]++
	}

	// Five of a kind or Four of a kind
	for _, val := range labels {
		if val == 5 {
			return "fiveKind"
		}
		if val == 4 {
			return "fourKind"
		}
	}

	// Full house
	appeared2 := false
	appeared3 := false
	for _, val := range labels {
		if val == 3 {
			if appeared2 {
				return "fullHouse"
			} else {
				appeared3 = true
			}
		}
		if val == 2 {
			if appeared3 {
				return "fullHouse"
			} else {
				appeared2 = true
			}
		}
	}

	// Three of a kind
	for _, val := range labels {
		if val == 3 {
			return "threeKind"
		}
	}

	// Two pair
	twoCounter := 0
	for _, val := range labels {
		if val == 2 {
			twoCounter++
		}

		if twoCounter == 2 {
			return "twoPair"
		}
	}

	// High card
	highCard := true
	for _, val := range labels {
		if val != 1 {
			highCard = false
		}
	}
	if highCard {
		return "highCard"
	}

	return "onePair"
}

/*
1. Wczytac wszystkie linie
2. Nalozyc na kazda wartosc jeden z typow - rozdzielic je na kategorie (osobne struktury)
3. Posortowac kategorie osobno
4. Na koniec zliczyc
*/
