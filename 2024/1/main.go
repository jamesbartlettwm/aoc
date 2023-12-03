package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed actual.txt
var actual []byte

func main() {
	partOneResult := partOne(actual)
	fmt.Printf("Part one: %d\n", partOneResult)

	partTwoResult := partTwo(actual)
	fmt.Printf("Part two: %d\n", partTwoResult)
}

func partOne(data []byte) int {
	lines := strings.Split(string(data), "\n")

	sum := 0
	for _, line := range lines {
		leftIndex := strings.IndexFunc(line, func(r rune) bool { return r >= '0' && r <= '9' })
		rightIndex := strings.LastIndexFunc(line, func(r rune) bool { return r >= '0' && r <= '9' })
		value, _ := strconv.Atoi(string(line[leftIndex]) + string(line[rightIndex]))
		sum += value
	}

	return sum
}

func partTwo(data []byte) int {
	lines := strings.Split(string(data), "\n")
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	sum := 0
	for _, line := range lines {
		leftI := len(line)
		rightI := 0
		left := 0
		right := 0
		for i, number := range numbers {
			index := strings.Index(line, number)
			rightIndex := strings.LastIndex(line, number)
			if index != -1 {
				if index <= leftI {
					leftI = index
					left = i
				}
			}

			if rightIndex != -1 {
				if rightIndex >= rightI {
					rightI = rightIndex
					right = i
				}
			}
		}
		ls := strconv.Itoa((left % 9) + 1)
		rs := strconv.Itoa((right % 9) + 1)
		value, _ := strconv.Atoi(ls + rs)
		sum += value
	}

	return sum
}
