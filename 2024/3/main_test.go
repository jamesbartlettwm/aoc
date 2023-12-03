package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample []byte

func TestPartOne(t *testing.T) {
	const expected = 4361
	result := partOne(sample)
	assert.Equal(t, expected, result)
}

func TestPartTwo(t *testing.T) {
	const expected = 467835
	result := partTwo(sample)
	assert.Equal(t, expected, result)
}
