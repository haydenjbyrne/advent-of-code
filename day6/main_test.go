package main

import (
	_ "embed"
	"testing"
)

//go:embed testInput.txt
var testInput string

func TestPart1(t *testing.T) {
	result := part1(testInput)
	if result != 288 {
		t.Fatalf(`part1(testInput) = %d, want 288`, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	if result != 71503 {
		t.Fatalf(`part2(testInput) = %d, want 71503`, result)
	}
}
