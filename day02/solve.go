package day02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/yokall/aoc-2025/utils"
)

func SolveDay02(input string) (int, int) {
	part1 := solvePart1(input)
	part2 := solvePart2(input)
	return part1, part2
}

func solvePart1(input string) int {
	lines := utils.ReadLinesString(input)

	sum := 0

	ranges := strings.Split(lines[0], ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidId(i) {
				sum += i
			}
		}
	}

	return sum
}

func isInvalidId(id int) bool {
	str := fmt.Sprintf("%d", id)
	length := len(str)

	if length%2 != 0 {
		return false
	}

	half := length / 2

	firstHalf := str[:half]
	secondHalf := str[half:]

	return firstHalf == secondHalf
}

func solvePart2(input string) int {
	lines := utils.ReadLinesString(input)

	sum := 0

	ranges := strings.Split(lines[0], ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidId2(i) {
				sum += i
			}
		}
	}

	return sum
}

func isInvalidId2(id int) bool {
	str := fmt.Sprintf("%d", id)
	length := len(str)
	half := length / 2

	for cs := 1; cs <= half; cs++ {
		invalid := true
		firstChunk := ""
		for chunk := range slices.Chunk(strings.Split(str, ""), cs) {
			if firstChunk == "" {
				firstChunk = strings.Join(chunk, "")
			} else {
				if strings.Join(chunk, "") != firstChunk {
					invalid = false
					break
				}
			}
		}

		if invalid {
			return true
		}
	}

	return false
}
