package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	log.Println("Reading input...")
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumCalibrationValues := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sumCalibrationValues += 10*getFirstNumber(line) + getLastNumber(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of calibration values = %d\n", sumCalibrationValues)
}

func getFirstNumber(line string) int {
	for i := range line {
		isInteger, integer := isNumberAtIndex(line, i)
		if isInteger {
			return integer
		}
	}

	log.Fatal("No number in line")
	return 0
}

func getLastNumber(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		isInteger, integer := isNumberAtIndex(line, i)
		if isInteger {
			return integer
		}
	}

	log.Fatal("No number in line")
	return 0
}

func isNumberAtIndex(line string, index int) (isNumber bool, number int) {
	c := rune(line[index])
	if unicode.IsDigit(c) {
		return true, int(c - '0')
	}

	for numberString, number := range numbers {
		if containsAtIndex(line, numberString, index) {
			return true, number
		}
	}

	return false, 0
}

func containsAtIndex(string, substring string, index int) bool {
	if index+len(substring) > len(string) {
		return false
	}

	for i := range substring {
		if substring[i] != string[index+i] {
			return false
		}
	}

	return true
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
