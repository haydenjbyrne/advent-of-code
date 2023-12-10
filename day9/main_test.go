package main

import (
	_ "embed"
	"testing"
)

//go:embed testInput.txt
var testInput string

func TestPart1(t *testing.T) {
	want := 114
	result := part1(testInput)
	if result != want {
		t.Fatalf(`part1(testInput) = %d, want %d`, result, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2
	result := part2(testInput)
	if result != want {
		t.Fatalf(`part2(testInput) = %d, want %d`, result, want)
	}
}
