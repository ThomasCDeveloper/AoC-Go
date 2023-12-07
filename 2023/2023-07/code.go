package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

type Bet struct {
	cards []int
	bet   int
	i     int
}

func Compare(bet1, bet2 Bet) bool {
	return bet1.cards[0] > bet2.cards[0]
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
		currentBet := Bet{cards, bet, i}

		bets = append(bets, currentBet)
	}

	wins := make([]int, len(data))
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if Compare(bets[i], bets[j]) {
				wins[i]++
			} else {
				wins[j]++
			}
		}
	}
	// sort bets according to wins

	return 0
}

func SolvePart2(data []string) int {
	return 0
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")
	fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
