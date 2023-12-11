package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	return sumDistances(input, 1)
}

func part2(input string) int {
	return sumDistances(input, 1000000)
}

func sumDistances(input string, expansionRate int) int {
	space := parseSpace(input)
	emptyRows := getEmptyRows(space)
	emptyCols := getEmptyColumns(space)

	sumDistances := 0
	galaxies := getGalaxies(space)
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			a := galaxies[i]
			b := galaxies[j]
			sumDistances += distance(a, b) + countInRange(emptyCols, a.col, b.col)*(expansionRate-1) + countInRange(emptyRows, a.row, b.row)*(expansionRate-1)
		}
	}
	return sumDistances
}

func parseSpace(input string) (grid [][]rune) {
	for _, rawRow := range strings.Split(input, "\n") {
		grid = append(grid, []rune(rawRow))
	}
	return
}

func getEmptyRows(galaxies [][]rune) (empty []int) {
	for i := range galaxies {
		if rowEmpty(galaxies, i) {
			empty = append(empty, i)
		}
	}
	fmt.Printf("Empty rows: %v\n", empty)
	return
}

func rowEmpty(galaxies [][]rune, row int) bool {
	for col := 0; col < len(galaxies[row]); col++ {
		if galaxies[row][col] == '#' {
			return false
		}
	}
	return true
}

func getEmptyColumns(galaxies [][]rune) (empty []int) {
	// assumption: all rows same length
	for i := range galaxies[0] {
		if columnEmpty(galaxies, i) {
			empty = append(empty, i)
		}
	}
	fmt.Printf("Empty cols: %v\n", empty)
	return
}

func columnEmpty(galaxies [][]rune, col int) bool {
	for row := 0; row < len(galaxies[0]); row++ {
		if galaxies[row][col] == '#' {
			return false
		}
	}
	return true
}

func getGalaxies(space [][]rune) (galaxies []position) {
	for rowIndex, row := range space {
		for colIndex, x := range row {
			if x == '#' {
				galaxies = append(galaxies, position{rowIndex, colIndex})
			}
		}
	}
	return
}

func countInRange(arr []int, x, y int) (count int) {
	start := min(x, y)
	end := max(x, y)
	for _, i := range arr {
		if i > start && i < end {
			count++
		}
	}
	return
}

type position struct{ row, col int }

func distance(a, b position) int {
	return max(a.row, b.row) - min(a.row, b.row) + max(a.col, b.col) - min(a.col, b.col)
}
