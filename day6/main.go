package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	match := regexp.MustCompile(`Time: (.*)\nDistance: (.*)`).FindStringSubmatch(input)
	times := parseNumbers(match[1])
	distances := parseNumbers(match[2])

	marginOfError := 1
	for i := 0; i < len(times); i++ {
		marginOfError *= calculateWaysToWin(times[i], distances[i])
	}
	return marginOfError
}

func part2(input string) int {
	match := regexp.MustCompile(`Time: (.*)\nDistance: (.*)`).FindStringSubmatch(input)
	time := parseNumberIgnoreWhitespace(match[1])
	distance := parseNumberIgnoreWhitespace(match[2])

	return calculateWaysToWin(time, distance)
}

func parseNumberIgnoreWhitespace(s string) int {
	timeStr := strings.Join(strings.Fields(s), "")
	time, timeErr := strconv.Atoi(timeStr)
	if timeErr != nil {
		log.Fatal(timeErr)
	}
	return time
}

func calculateWaysToWin(raceTime, bestDistance int) int {
	count := 0
	for i := 0; i < raceTime; i++ {
		distanceTraveled := i * (raceTime - i)
		if distanceTraveled > bestDistance {
			count++
		}
	}
	return count
}

func parseNumbers(s string) []int {
	var numbers []int
	rawNumbers := strings.Fields(s)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}
