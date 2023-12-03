package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample []byte

func TestPartOne(t *testing.T) {
	const expected = 8
	result := partOne(sample)
	assert.Equal(t, expected, result)
}

func TestPartTwo(t *testing.T) {
	const expected = 2286
	result := partTwo(sample)
	assert.Equal(t, expected, result)
}
