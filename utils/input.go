package utils

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadLinesString reads input as a string and returns lines as a slice of strings
func ReadLinesString(input string) []string {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadRawFile reads entire file as a single string
func ReadRawFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Read2DGrid converts lines into a 2D character grid
// Each line becomes a row, each character a column
func Read2DGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// Read2DGridFromString reads input string into a 2D character grid
func Read2DGridFromString(input string) [][]rune {
	lines := ReadLinesString(input)
	return Read2DGrid(lines)
}

// ParseIntLines converts lines of single integers to []int
func ParseIntLines(lines []string) []int {
	var result []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

// FilterEmptyLines removes empty strings from slice
func FilterEmptyLines(lines []string) []string {
	var filtered []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

// SplitOn splits lines by a delimiter string
func SplitOn(lines []string, delimiter string) [][]string {
	var groups [][]string
	var current []string

	for _, line := range lines {
		if line == delimiter {
			if len(current) > 0 {
				groups = append(groups, current)
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}

	if len(current) > 0 {
		groups = append(groups, current)
	}

	if groups == nil {
		groups = [][]string{}
	}
	return groups
}
