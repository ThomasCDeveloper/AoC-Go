package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Bet struct {
	cards []int
	bet   int
	i     int
	wins  int
	rep   []int
	J     int
}

func GetCardRepartition(cards []int) []int {
	output := make([]int, 13)
	for i := 0; i < 5; i++ {
		output[cards[i]]++
	}
	return output
}

func Compare(bet1, bet2 Bet) bool {
	cards1 := bet1.cards
	cards2 := bet2.cards
	val1 := slices.Max(bet1.rep) + bet1.J
	val2 := slices.Max(bet2.rep) + bet2.J

	if val1 > val2 {
		return true
	}
	if val2 > val1 {
		return false
	}

	// check for 3-2 configurations
	if val1 == 3 && val2 == 3 {
		if slices.Contains(bet1.rep, 2) && !slices.Contains(bet2.rep, 2) {
			return true
		}
		if !slices.Contains(bet1.rep, 2) && slices.Contains(bet2.rep, 2) {
			return false
		}
	}
	// add condition HERE
	// si 1 double + 1 joker ça bat 1 double
	if bet1.J == 0 && bet2.J == 0 {
		if val1 == 2 && val2 == 2 {
			// change that to count 2s in reps
			count1, count2 := 0, 0
			for i := 0; i < len(bet1.rep); i++ {
				if bet1.rep[i] == 2 {
					count1++
				}
				if bet2.rep[i] == 2 {
					count2++
				}
			}
			if count1 > count2 {
				return true
			}
			if count2 > count1 {
				return false
			}
		}
	}

	for i := 0; i < 5; i++ {
		if cards1[i] > cards2[i] {
			return true
		}
		if cards2[i] > cards1[i] {
			return false
		}
	}
	return true
}

func SolvePart1(data []string) int {
	bets := []Bet{}
	for i, line := range data {
		cards := []int{}
		deck := strings.Split(line, " ")[0]
		for i := 0; i < 5; i++ {
			cards = append(cards, strings.Index("23456789TJQKA", string(deck[i])))
		}
		bet, _ := strconv.Atoi(strings.Split(line, " ")[1])
		currentBet := Bet{cards, bet, i, 0, GetCardRepartition(cards), 0}

		bets = append(bets, currentBet)
	}

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if Compare(bets[i], bets[j]) {
				bets[i].wins++
			} else {
				bets[j].wins++
			}
		}
	}

	// sort bets according to wins
	sort.Slice(bets, func(i, j int) bool {
		return bets[i].wins < bets[j].wins
	})

	sum := 0

	for i, bet := range bets {
		sum += (i + 1) * bet.bet
	}

	return sum
}

func SolvePart2(data []string) int {
	return 0
}

func main() {
	data := GetInput("test.txt")

	// PART 1
	fmt.Println("Part 1:")
	fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
