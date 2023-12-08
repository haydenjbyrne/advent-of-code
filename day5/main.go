package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
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
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	seeds := parseSeeds(parts[0])
	almanac := almanac{maps: parseMaps(parts[1:])}

	lowestLocationNumber := math.MaxInt
	for _, seed := range seeds {
		locationNumber := almanac.mapLocation(seed, "seed", "location")
		if locationNumber < lowestLocationNumber {
			lowestLocationNumber = locationNumber
		}
	}
	return lowestLocationNumber
}

func part2(input string) int {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	seedRanges := parseSeedRanges(parts[0])
	almanac := almanac{maps: parseMaps(parts[1:])}

	locationRanges := almanac.mapRanges(seedRanges, "seed", "location")
	lowestLocationNumber := math.MaxInt
	for _, locationRange := range locationRanges {
		if locationRange.start < lowestLocationNumber {
			lowestLocationNumber = locationRange.start
		}
	}
	return lowestLocationNumber
}

func parseSeeds(rawSeeds string) []int {
	match := regexp.MustCompile(`seeds: (.*)`).FindStringSubmatch(rawSeeds)
	return parseNumbers(match[1])
}

func parseSeedRanges(rawSeeds string) []Range {
	match := regexp.MustCompile(`seeds: (.*)`).FindStringSubmatch(rawSeeds)
	numbers := parseNumbers(match[1])
	var seeds []Range
	for i := 0; i < len(numbers); i += 2 {
		seeds = append(seeds, Range{start: numbers[i], length: numbers[i+1]})
	}
	return seeds
}

func parseMaps(mapsRaw []string) []categoryMap {
	var maps []categoryMap
	for _, mapRaw := range mapsRaw {
		maps = append(maps, parseMap(mapRaw))
	}
	return maps
}

func parseMap(mapRaw string) categoryMap {
	parts := strings.Split(mapRaw, "\n")
	match := regexp.MustCompile(`(.*)-to-(.*) map:`).FindStringSubmatch(parts[0])
	return categoryMap{source: match[1], destination: match[2], ranges: parseRangeMaps(parts[1:])}
}

func parseRangeMaps(rangesRaw []string) []rangeMap {
	var ranges []rangeMap
	for _, rangeRaw := range rangesRaw {
		numbers := parseNumbers(rangeRaw)
		ranges = append(ranges, rangeMap{source: Range{start: numbers[1], length: numbers[2]}, destinationStart: numbers[0]})
	}
	return ranges
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

type almanac struct {
	maps []categoryMap
}

func (a almanac) mapLocation(num int, sourceCategory, destinationCategory string) int {
	if sourceCategory == destinationCategory {
		return num
	}

	for _, mp := range a.maps {
		if mp.source == sourceCategory {
			return a.mapLocation(mp.mapNumber(num), mp.destination, destinationCategory)
		}
	}

	log.Fatalf("No mapping from %s to %s", sourceCategory, destinationCategory)
	return 0
}

func (a almanac) mapRanges(seed []Range, sourceCategory, destinationCategory string) []Range {
	if sourceCategory == destinationCategory {
		return seed
	}

	for _, mp := range a.maps {
		if mp.source == sourceCategory {
			fmt.Printf("Mapping from %s to %s\n", mp.source, mp.destination)
			return a.mapRanges(mp.mapRanges(seed), mp.destination, destinationCategory)
		}
	}

	log.Fatalf("No mapping from %s to %s", sourceCategory, destinationCategory)
	return []Range{}
}

type categoryMap struct {
	destination string
	source      string
	ranges      []rangeMap
}

func (mp categoryMap) mapRanges(ranges []Range) []Range {
	var mappedRanges []Range
	for _, r := range ranges {
		mappedRanges = append(mappedRanges, mp.mapRange(r)...)
	}
	return mappedRanges
}

func (mp categoryMap) mapRange(source Range) []Range {
	fmt.Printf("Mapping range: %v\n", source)
	for _, rangeMap := range mp.ranges {
		intersection := source.intersect(rangeMap.source)
		if intersection.length > 0 {
			mappedRange := Range{start: intersection.start - rangeMap.source.start + rangeMap.destinationStart, length: intersection.length}
			fmt.Printf("Mapped %v to %v\n", intersection, mappedRange)
			excludedRanges := source.exclude(intersection)
			return append(mp.mapRanges(excludedRanges), mappedRange)
		}
	}
	// no mapping
	fmt.Printf("Unmapped range: %v\n", source)
	return []Range{source}
}

func (mp categoryMap) mapNumber(num int) int {
	for _, mappedRange := range mp.ranges {
		if num >= mappedRange.source.start && num <= mappedRange.source.End() {
			return num - mappedRange.source.start + mappedRange.destinationStart
		}
	}
	// no mapping
	return num
}

type Range struct {
	start  int
	length int
}

func (r Range) End() int {
	return r.start + r.length - 1
}

func (r Range) intersect(compare Range) Range {
	start := max(r.start, compare.start)
	end := min(r.End(), compare.End())
	if start > end {
		return Range{start: 0, length: 0}
	}
	return Range{start: start, length: end - start + 1}
}

func (r Range) exclude(exclude Range) []Range {
	var ranges []Range
	if r.start < exclude.start {
		ranges = append(ranges, Range{start: r.start, length: exclude.start - r.start})
	}
	if r.End() > exclude.End() {
		ranges = append(ranges, Range{start: exclude.End() + 1, length: r.End() - exclude.End()})
	}
	return ranges
}

type rangeMap struct {
	source           Range
	destinationStart int
}
