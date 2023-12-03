package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed actual.txt
var actual []byte

type Set = map[string]int

type Game struct {
	id   int
	sets []Set
}

func main() {
	partOneResult := partOne(actual)
	fmt.Printf("Part one: %d\n", partOneResult)

	partTwoResult := partTwo(actual)
	fmt.Printf("Part two: %d\n", partTwoResult)
}

func partOne(data []byte) int {
	lines := strings.Split(string(data), "\n")

	maximums := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := make([]*Game, len(lines))
	for i, line := range lines {
		games[i] = parseLine(line)
	}

	sum := 0
	for _, game := range games {
		if checkGameValidity(game, maximums) {
			sum += game.id
		}
	}

	return sum
}

func partTwo(data []byte) int {
	lines := strings.Split(string(data), "\n")

	games := make([]*Game, len(lines))
	for i, line := range lines {
		games[i] = parseLine(line)
	}

	sum := int(0)

	for _, game := range games {
		maxGreen := slices.MaxFunc(game.sets, func(a map[string]int, b map[string]int) int {
			return cmp.Compare(a["green"], b["green"])
		})["green"]
		maxBlue := slices.MaxFunc(game.sets, func(a map[string]int, b map[string]int) int {
			return cmp.Compare(a["blue"], b["blue"])
		})["blue"]
		maxRed := slices.MaxFunc(game.sets, func(a map[string]int, b map[string]int) int {
			return cmp.Compare(a["red"], b["red"])
		})["red"]

		power := maxGreen * maxBlue * maxRed
		sum += power
	}

	return sum
}

func checkGameValidity(game *Game, maximums map[string]int) bool {
	for _, set := range game.sets {
		for color, count := range set {
			if maximums[color] < count {
				return false
			}
		}
	}

	return true
}

func parseLine(line string) *Game {
	idAndData := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Replace(idAndData[0], "Game ", "", -1))

	game := &Game{id: id, sets: []Set{}}

	cubeSets := strings.Split(idAndData[1], "; ")
	for _, setData := range cubeSets {
		set := map[string]int{}
		cubeData := strings.Split(setData, ", ")

		for _, cubes := range cubeData {
			numColor := strings.Split(cubes, " ")
			value, _ := strconv.Atoi(numColor[0])
			set[numColor[1]] = int(value)
		}
		game.sets = append(game.sets, set)
	}

	return game
}
