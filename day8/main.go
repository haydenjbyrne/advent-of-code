package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strings"
)

//go:embed part1.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	instructions := strings.Split(parts[0], "")
	nodePaths := parseNodePaths(parts[1])
	nodePathMap := map[string]nodePath{}
	for _, path := range nodePaths {
		nodePathMap[path.from] = path
	}

	count := 0
	position := "AAA"
	for ; position != "ZZZ"; count++ {
		if count > 1_000_000 {
			log.Fatal("Suspected infinite loop")
		}

		fmt.Printf("Position: %s\n", position)
		instruction := instructions[count%len(instructions)]
		nodePath := nodePathMap[position]
		position = nodePath.to[instruction]
	}

	return count
}

func part2(input string) int {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	instructions := strings.Split(parts[0], "")
	nodePaths := parseNodePaths(parts[1])
	nodePathMap := map[string]nodePath{}
	for _, path := range nodePaths {
		nodePathMap[path.from] = path
	}

	var positions []string
	for _, path := range nodePaths {
		if path.from[len(path.from)-1] == 'A' {
			positions = append(positions, path.from)
		}
	}

	fmt.Printf("Instruction Count: %d\n", len(instructions))

	var counts []int
	for _, startPosition := range positions {
		count, _ := getNextEndPosition(startPosition, instructions, nodePathMap)
		counts = append(counts, count)
	}

	// observed all routes get into separate cycles - the point they all meet is the lowest common multiple
	// this may not work with all data sets.
	return leastCommonMultiple(counts)
}

func leastCommonMultiple(s []int) int {
	if len(s) == 0 {
		return 0
	}

	lcm := 1
	primeFactorial := getPrimeFactors(s[0])
	for _, prime := range primeFactorial {
		if allDivisibleBy(s, prime) {
			s = divideAllBy(s, prime)
			lcm *= prime
		}
	}

	for _, i := range s {
		lcm *= i
	}
	return lcm
}

func getPrimeFactors(n int) (pfs []int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			// extract prime factor
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// n now prime
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func allDivisibleBy(ns []int, x int) bool {
	for _, n := range ns {
		if n%x != 0 {
			return false
		}
	}
	return true
}

func divideAllBy(ns []int, x int) (divided []int) {
	for _, n := range ns {
		divided = append(divided, n/x)
	}
	return
}

func getNextEndPosition(startPosition string, instructions []string, nodePathMap map[string]nodePath) (int, string) {
	count := 0
	position := startPosition
	for {
		if count > 1_000_000 {
			log.Fatal("Suspected infinite loop")
		}

		instruction := instructions[count%len(instructions)]
		nodePath := nodePathMap[position]
		position = nodePath.to[instruction]
		count++
		if isEndPosition(position) {
			break
		}
	}
	fmt.Printf("Position: %s - %s (%d/%d=%f)\n", startPosition, position, count, len(instructions), float64(count)/float64(len(instructions)))
	return count, position
}

func isEndPosition(position string) bool {
	return position[len(position)-1] == 'Z'
}

type nodePath struct {
	from string
	to   map[string]string
}

var nodePathPattern = regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

func parseNodePaths(s string) []nodePath {
	rawNodePaths := strings.Split(s, "\n")
	var nodePaths []nodePath
	for _, rawNodePath := range rawNodePaths {
		match := nodePathPattern.FindStringSubmatch(rawNodePath)
		path := nodePath{
			from: match[1],
			to: map[string]string{
				"L": match[2],
				"R": match[3],
			}}
		nodePaths = append(nodePaths, path)
	}
	return nodePaths
}
