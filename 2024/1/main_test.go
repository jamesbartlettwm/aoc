package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample []byte

//go:embed sample2.txt
var sample2 []byte

func TestPartOne(t *testing.T) {
	const expected = 142
	result := partOne(sample)
	assert.Equal(t, expected, result)
}

func TestPartTwo(t *testing.T) {
	const expected = 281
	result := partTwo(sample2)
	assert.Equal(t, expected, result)
}
