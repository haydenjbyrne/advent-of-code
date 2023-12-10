package main

import (
	"adventOfCode/ParseUtils"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	sequences := parseSequences(input)

	sumNextSequenceNumbers := 0
	for _, sequence := range sequences {
		sumNextSequenceNumbers += getNextSequenceNumber(sequence)
	}
	return sumNextSequenceNumbers
}

func part2(input string) int {
	sequences := parseSequences(input)

	sumNextSequenceNumbers := 0
	for _, sequence := range sequences {
		slices.Reverse(sequence)
		sumNextSequenceNumbers += getNextSequenceNumber(sequence)
	}
	return sumNextSequenceNumbers
}

func parseSequences(input string) (sequences [][]int) {
	for _, rawSequence := range strings.Split(input, "\n") {
		if len(rawSequence) == 0 {
			continue
		}
		sequences = append(sequences, ParseUtils.ParseNumbers(rawSequence))
	}
	return
}

func getNextSequenceNumber(sequence []int) int {
	fmt.Printf("Sequence: %v\n", sequence)

	if allElementsZero(sequence) {
		return 0
	}

	var nextSequence []int
	for i := 0; i < len(sequence)-1; i++ {
		nextSequence = append(nextSequence, sequence[i+1]-sequence[i])
	}
	return sequence[len(sequence)-1] + getNextSequenceNumber(nextSequence)
}

func allElementsZero(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			return false
		}
	}
	return true
}
