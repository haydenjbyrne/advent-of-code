package main

import (
	_ "embed"
	"testing"
)

//go:embed testInput/part1.txt
var testInput string

func TestPart1(t *testing.T) {
	want := 374
	result := part1(testInput)
	if result != want {
		t.Fatalf(`part1(testInput) = %d, want %d`, result, want)
	}
}
