package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
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
	cards := parseCards(input)

	totalPoints := 0
	for _, card := range cards {
		totalPoints += card.calculateCardValue()
	}
	return totalPoints
}

func part2(input string) int {
	cards := parseCards(input)

	countOfEachCard := make([]int, len(cards))
	for i, card := range cards {
		countOfEachCard[i] += 1
		for j := i + 1; j <= i+card.countMatchingNumbers(); j++ {
			countOfEachCard[j] += countOfEachCard[i]
		}
	}

	sumCards := 0
	for _, countOfCard := range countOfEachCard {
		sumCards += countOfCard
	}
	return sumCards
}

func (c card) calculateCardValue() int {
	matchingNumbers := c.countMatchingNumbers()
	if matchingNumbers == 0 {
		return 0
	}
	return 1 << (matchingNumbers - 1)
}

func (c card) countMatchingNumbers() int {
	matchingNumbers := 0
	for _, number := range c.numbers {
		for _, winningNumber := range c.winningNumbers {
			if number == winningNumber {
				matchingNumbers++
			}
		}
	}
	return matchingNumbers
}

func parseCards(input string) []card {
	var cards []card
	rawCards := strings.Split(input, "\n")
	for _, rawCard := range rawCards {
		if rawCard == "" {
			continue
		}
		cards = append(cards, parseCard(rawCard))
	}
	return cards
}

func parseCard(rawCard string) card {
	match := regexp.MustCompile(`Card\s*(\d*): (.*) \| (.*)`).FindStringSubmatch(rawCard)
	if match == nil {
		log.Fatalf("Invalid card format: %s", rawCard)
	}
	return card{winningNumbers: parseNumbers(match[2]), numbers: parseNumbers(match[3])}
}

func parseNumbers(s string) []int {
	var numbers []int
	rawNumbers := strings.Fields(s)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

type card struct {
	winningNumbers []int
	numbers        []int
}
