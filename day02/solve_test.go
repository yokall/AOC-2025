package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	part1, part2 := SolveDay02(input)

	assert.Equal(t, 1227775554, part1)
	assert.Equal(t, 4174379265, part2)
}

func TestIsInvalidId2(t *testing.T) {
	assert.Equal(t, true, isInvalidId2(11))
	assert.Equal(t, true, isInvalidId2(22))
	assert.Equal(t, true, isInvalidId2(99))
	assert.Equal(t, true, isInvalidId2(1010))
	assert.Equal(t, true, isInvalidId2(1188511885))
	assert.Equal(t, true, isInvalidId2(222222))
	assert.Equal(t, true, isInvalidId2(446446))
	assert.Equal(t, true, isInvalidId2(38593859))
	assert.Equal(t, true, isInvalidId2(1212121212))
}
