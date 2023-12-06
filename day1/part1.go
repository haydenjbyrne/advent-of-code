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

	scanner := bufio.NewScanner(file)

	sumCalibrationValues := 0

	for scanner.Scan() {
		line := scanner.Text()
		sumCalibrationValues += 10*firstDigit(line) + lastDigit(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of calibration values = %d\n", sumCalibrationValues)
}

func firstDigit(line string) int {
	for _, c := range line {
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return 0
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return 0
}
