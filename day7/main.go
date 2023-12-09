package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	hands := parseHands(input)

	return calculateWinnings(hands, false)
}

func part2(input string) int {
	hands := parseHands(input)

	return calculateWinnings(hands, true)
}

func calculateWinnings(hands []hand, jokers bool) int {
	rankHands(hands, jokers)

	winnings := 0
	for i, h := range hands {
		handRank := i + 1
		log.Printf("%v: %v : %v", handRank, string(h.cards), h.getTypeRank(jokers))
		winnings += h.bid * handRank
	}
	return winnings
}

func rankHands(hands []hand, jokers bool) {
	sort.Slice(hands, func(i, j int) bool {
		// compare by hand type
		rankI := hands[i].getTypeRank(jokers)
		rankJ := hands[j].getTypeRank(jokers)
		if rankI != rankJ {
			return rankI < rankJ
		}

		// then by card rank
		for x := range hands[i].cards {
			if hands[i].cards[x] == hands[j].cards[x] {
				continue
			}
			return getCardRank(hands[i].cards[x], jokers) < getCardRank(hands[j].cards[x], jokers)
		}

		return false
	})
}

func parseHands(input string) []hand {
	rawBids := strings.Split(strings.TrimRight(input, "\n"), "\n")
	var hands []hand
	for _, rawBid := range rawBids {
		hands = append(hands, parseHand(rawBid))
	}
	return hands
}

func parseHand(rawHand string) hand {
	parts := strings.Fields(rawHand)
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return hand{[]rune(parts[0]), bid}
}

func (h hand) getTypeRank(jokers bool) int {
	cardCounts := h.getCardCounts()

	if jokers {
		swapOutJokers(cardCounts)
	}

	return calculateHandKind(cardCounts)
}

func (h hand) getCardCounts() map[rune]int {
	cardCounts := map[rune]int{}
	for _, card := range h.cards {
		cardCounts[card]++
	}
	return cardCounts
}

func swapOutJokers(cardCounts map[rune]int) {
	maxCardCount := 0
	var maxCard rune
	for card, count := range cardCounts {
		if count > maxCardCount && card != 'J' {
			maxCardCount = count
			maxCard = card
		}
	}
	if maxCardCount > 0 {
		cardCounts[maxCard] += cardCounts['J']
		delete(cardCounts, 'J')
	}
}

func calculateHandKind(cardCounts map[rune]int) int {
	if len(cardCounts) == 1 {
		return fiveOfAKind
	}
	if len(cardCounts) == 2 {
		for _, count := range cardCounts {
			if count == 4 {
				return fourOfAKind
			}
		}
		return fullHouse
	}
	if len(cardCounts) == 3 {
		for _, count := range cardCounts {
			if count == 3 {
				return threeOfAKind
			}
		}
		return twoPair
	}
	if len(cardCounts) == 4 {
		return onePair
	}
	return highCard
}

type hand struct {
	cards []rune
	bid   int
}

var cardRanks = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
}

func getCardRank(c rune, jokers bool) int {
	if jokers && c == 'J' {
		return 1
	}
	mappedRank, ok := cardRanks[c]
	if ok {
		return mappedRank
	}
	return int(c - '0')
}

const (
	fiveOfAKind  = 6
	fourOfAKind  = 5
	fullHouse    = 4
	threeOfAKind = 3
	twoPair      = 2
	onePair      = 1
	highCard     = 0
)
