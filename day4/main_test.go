package main

import (
	_ "embed"
	"testing"
)

//go:embed testInput.txt
var testInput string

func TestPart1(t *testing.T) {
	if part1(testInput) != 13 {
		t.Fail()
	}
}

func TestCard1(t *testing.T) {
	if part1("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53") != 8 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	if result != 30 {
		t.Fail()
		t.Fatalf(`part2(testInput) = %d, want 30`, result)
	}
}
