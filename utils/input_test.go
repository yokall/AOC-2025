package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLinesString(t *testing.T) {
	input := "line1\nline2\nline3"
	expected := []string{"line1", "line2", "line3"}

	result := ReadLinesString(input)

	assert.Equal(t, expected, result)
}

func TestReadLinesStringEmpty(t *testing.T) {
	input := ""
	var expected []string

	result := ReadLinesString(input)

	assert.Equal(t, expected, result)
}

func TestRead2DGrid(t *testing.T) {
	lines := []string{"abc", "def", "ghi"}
	expected := [][]rune{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}

	result := Read2DGrid(lines)

	assert.Equal(t, expected, result)
}

func TestRead2DGridFromString(t *testing.T) {
	input := "abc\ndef\nghi"
	expected := [][]rune{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}

	result := Read2DGridFromString(input)

	assert.Equal(t, expected, result)
}

func TestRead2DGridEmpty(t *testing.T) {
	lines := []string{}
	expected := [][]rune{}

	result := Read2DGrid(lines)

	assert.Equal(t, expected, result)
}

func TestFilterEmptyLines(t *testing.T) {
	input := []string{"line1", "", "line2", "  ", "line3"}
	expected := []string{"line1", "line2", "line3"}

	result := FilterEmptyLines(input)

	assert.Equal(t, expected, result)
}

func TestSplitOn(t *testing.T) {
	input := []string{"a", "b", "---", "c", "d", "---", "e"}
	expected := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e"},
	}

	result := SplitOn(input, "---")

	assert.Equal(t, expected, result)
}

func TestSplitOnEmpty(t *testing.T) {
	input := []string{}
	expected := [][]string{}

	result := SplitOn(input, "---")

	assert.Equal(t, expected, result)
}

func TestSplitOnNoDelimiter(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := [][]string{
		{"a", "b", "c"},
	}

	result := SplitOn(input, "---")

	assert.Equal(t, expected, result)
}
