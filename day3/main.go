package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	numbers := parseNumbers(rows)

	part1(numbers, rows)

	part2(rows, numbers)
}

func part1(numbers []numberPosition, rows []string) {
	sumPartNumbers := 0
	for _, number := range numbers {
		if isNumberNextToSymbol(number, rows) {
			sumPartNumbers += number.number
		}
	}
	fmt.Printf("Sum of part numbers = %d\n", sumPartNumbers)
}

func part2(rows []string, numbers []numberPosition) {
	sumGearRatios := 0
	for rowIndex, row := range rows {
		for colIndex, char := range row {
			isGear, gearRatio := isGear(rowIndex, colIndex, char, numbers)
			if isGear {
				sumGearRatios += gearRatio
			}
		}
	}
	fmt.Printf("Sum of gear ratios = %d\n", sumGearRatios)
}

func isGear(rowIndex, colIndex int, char rune, numbers []numberPosition) (isGear bool, gearRatio int) {
	if char != '*' {
		return false, 0
	}

	//could optimize but scanning all numbers is fast enough
	var adjacentNumbers []numberPosition
	for _, number := range numbers {
		if number.row-1 <= rowIndex && number.row+1 >= rowIndex && number.col-1 <= colIndex && number.col+number.len >= colIndex {
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}

	if len(adjacentNumbers) == 2 {
		return true, adjacentNumbers[0].number * adjacentNumbers[1].number
	}

	return false, 0
}

func parseNumbers(rows []string) []numberPosition {
	var numbers []numberPosition
	for rowIndex, row := range rows {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			isNumber, number, length := getNumberAtIndex(row, colIndex)
			if isNumber {
				numbers = append(numbers, numberPosition{number, position{col: colIndex, row: rowIndex, len: length}})
				// additional digits in number to skip
				colIndex += length - 1
			}
		}
	}
	return numbers
}

func getNumberAtIndex(line string, index int) (isNumber bool, number int, length int) {
	for ; index+length < len(line); length++ {
		c := rune(line[index+length])
		if !unicode.IsDigit(c) {
			break
		}
	}

	if length == 0 {
		return false, 0, length
	}

	num, err := strconv.ParseInt(line[index:index+length], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return true, int(num), length
}

func isNumberNextToSymbol(number numberPosition, rows []string) bool {
	minRow := max(number.row-1, 0)
	maxRow := min(number.row+1, len(rows)-1)
	minCol := max(number.col-1, 0)
	for row := minRow; row <= maxRow; row++ {
		maxCol := min(number.col+number.len, len(rows[row])-1)
		for col := minCol; col <= maxCol; col++ {
			c := rune(rows[row][col])
			if (unicode.IsSymbol(c) || unicode.IsPunct(c)) && c != '.' {
				return true
			}
		}
	}
	return false
}

type numberPosition struct {
	number int
	position
}

type position struct {
	row int
	col int
	len int
}
