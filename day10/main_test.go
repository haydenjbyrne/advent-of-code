package main

import (
	_ "embed"
	"testing"
)

//go:embed testInput/part1.txt
var testInput string

func TestPart1(t *testing.T) {
	want := 8
	result := part1(testInput)
	if result != want {
		t.Fatalf(`part1(testInput) = %d, want %d`, result, want)
	}
}

//go:embed testInput/part2-1.txt
var part2Input string

func TestPart2(t *testing.T) {
	want := 4
	result := part2(part2Input)
	if result != want {
		t.Fatalf(`part2(testInput) = %d, want %d`, result, want)
	}
}

//go:embed testInput/part2-2.txt
var part22Input string

func TestPart22(t *testing.T) {
	want := 10
	result := part2(part22Input)
	if result != want {
		t.Fatalf(`part2(testInput) = %d, want %d`, result, want)
	}
}
