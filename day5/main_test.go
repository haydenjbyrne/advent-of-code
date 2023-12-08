package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed testInput.txt
var testInput string

func TestPart1(t *testing.T) {
	result := part1(testInput)
	if result != 35 {
		t.Fatalf(`part1(testInput) = %d, want 35`, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	if result != 46 {
		t.Fatalf(`part2(testInput) = %d, want 46`, result)
	}
}

func TestIntersect(t *testing.T) {
	expected := Range{3, 2}
	result := Range{1, 4}.intersect(Range{3, 3})
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf(`intersect(testInput) = %v, want %v`, result, expected)
	}
}

func TestExclude(t *testing.T) {
	expected := []Range{{1, 2}}
	result := Range{1, 4}.exclude(Range{3, 2})
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf(`exclude(testInput) = %v, want %v`, result, expected)
	}
}
