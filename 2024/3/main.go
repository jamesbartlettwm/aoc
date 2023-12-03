package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	_ "embed"
)

//go:embed actual.txt
var actual []byte

type Number struct {
	value    int
	row      int
	startCol int
	endCol   int
}

type Symbol struct {
	row int
	col int
}

func main() {
	partOneResult := partOne(actual)
	fmt.Printf("Part one: %d\n", partOneResult)

	partTwoResult := partTwo(actual)
	fmt.Printf("Part two: %d\n", partTwoResult)
}

func partOne(data []byte) int {
	numbers, symbols := parse("", data)

	sum := 0
	for _, number := range numbers {
		for _, symbol := range symbols {
			if (symbol.col >= number.startCol-1 && symbol.col <= number.endCol+1) &&
				(symbol.row >= number.row-1 && symbol.row <= number.row+1) {
				sum += number.value
				break
			}
		}
	}

	return sum
}

func partTwo(data []byte) int {
	numbers, stars := parse("*", data)

	sum := 0
	for _, star := range stars {
		matchingNumbers := []Number{}
		for _, number := range numbers {
			if (star.col >= number.startCol-1 && star.col <= number.endCol+1) &&
				(star.row >= number.row-1 && star.row <= number.row+1) {
				matchingNumbers = append(matchingNumbers, number)
			}
		}
		if len(matchingNumbers) == 2 {
			sum += matchingNumbers[0].value * matchingNumbers[1].value
		}
	}

	return sum
}

func parse(validSymbols string, data []byte) ([]Number, []Symbol) {
	var s scanner.Scanner

	numbers := []Number{}
	symbols := []Symbol{}

	s.Init(strings.NewReader((string(data))))
	s.Mode = scanner.ScanInts | scanner.ScanChars

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		switch token {
		case scanner.Int:
			p := s.Position
			intValue, _ := strconv.Atoi(s.TokenText())
			numbers = append(numbers, Number{
				value:    intValue,
				row:      p.Line - 1,
				startCol: p.Column - 1,
				endCol:   p.Column + len(fmt.Sprint(intValue)) - 2,
			})
		default:
			if (validSymbols == "" && string(token) != ".") || strings.Contains(validSymbols, string(token)) {
				symbols = append(symbols, Symbol{
					row: s.Position.Line - 1,
					col: s.Position.Column - 1,
				})
			}
		}
	}

	return numbers, symbols
}
