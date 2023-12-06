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
	games := parseGames()

	part1(games)

	part2(games)
}

func part1(games []game) {
	sumPossibleGames := 0
	for _, game := range games {
		if gamePossible(game) {
			sumPossibleGames += game.id
		}
	}
	fmt.Printf("Sum of possible games = %d\n", sumPossibleGames)
}

func gamePossible(game game) bool {
	for _, round := range game.rounds {
		for colour, count := range round.cubeCounts {
			if cubeInventory[colour] < count {
				return false
			}
		}
	}
	return true
}

func part2(games []game) {
	sumGamePowers := 0
	for _, game := range games {
		minCubesRequired := map[string]int{}
		for _, round := range game.rounds {
			for colour, count := range round.cubeCounts {
				if count > minCubesRequired[colour] {
					minCubesRequired[colour] = count
				}
			}
		}
		power := 1
		for _, count := range minCubesRequired {
			power *= count
		}
		sumGamePowers += power
	}
	fmt.Printf("Sum of game powers = %d\n", sumGamePowers)
}

func parseGames() []game {
	lines := strings.Split(input, "\n")
	var games []game
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(rawGame string) game {
	match := regexp.MustCompile(`Game (\d*): (.*)`).FindStringSubmatch(rawGame)

	gameNumber, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatal(err)
	}

	rawRounds := strings.Split(match[2], "; ")
	var rounds []round
	for _, roundPart := range rawRounds {
		rounds = append(rounds, parseRound(roundPart))
	}
	return game{id: gameNumber, rounds: rounds}
}

func parseRound(rawRound string) round {
	regex := regexp.MustCompile(`(\d*) (\w*)`)
	rawCubeCounts := strings.Split(rawRound, ", ")
	cubeCounts := map[string]int{}
	for _, rawCubeCount := range rawCubeCounts {
		match := regex.FindStringSubmatch(rawCubeCount)
		if match == nil {
			log.Fatalf("Couldn't parse cube count %s", rawCubeCount)
		}
		cubeCount, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		cubeCounts[match[2]] = cubeCount
	}
	return round{cubeCounts: cubeCounts}
}

type game struct {
	id     int
	rounds []round
}

type round struct {
	cubeCounts map[string]int
}

var cubeInventory = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}
