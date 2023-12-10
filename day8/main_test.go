package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testInput/part1-1.txt
var part1Input1 string

func TestPart1(t *testing.T) {
	want := 2
	result := part1(part1Input1)
	if result != want {
		t.Fatalf(`part1(testInput) = %d, want %d`, result, want)
	}
}

//go:embed testInput/part1-2.txt
var part1input2 string

func TestPart1_2(t *testing.T) {
	want := 6
	result := part1(part1input2)
	if result != want {
		t.Fatalf(`part1(testInput) = %d, want %d`, result, want)
	}
}

//go:embed testInput/part2.txt
var part2Input string

func TestPart2(t *testing.T) {
	want := 6
	result := part2(part2Input)
	if result != want {
		t.Fatalf(`part2(testInput) = %d, want %d`, result, want)
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	want := 6
	result := leastCommonMultiple([]int{2, 3})
	if result != want {
		t.Fatalf(`leastCommonMultiple(testInput) = %d, want %d`, result, want)
	}

	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2, 3}, 6},
		{[]int{4, 6, 10}, 60},
		{[]int{7 * 13 * 13, 11 * 13 * 13, 19 * 13 * 13}, 7 * 11 * 19 * 13 * 13},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			ans := leastCommonMultiple(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
