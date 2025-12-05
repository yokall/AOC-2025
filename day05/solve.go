package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/yokall/aoc-2025/utils"
)

func SolveDay05(input string) (int, int) {
	lines := utils.ReadLinesString(input)
	groups := utils.SplitOn(lines, "")

	freshRanges := groups[0]
	ingredients := groups[1]

	part1 := solvePart1(ingredients, freshRanges)
	part2 := solvePart2(freshRanges)

	return part1, part2
}

func solvePart1(ingredients []string, freshRanges []string) int {
	count := 0

	for _, ingredient := range ingredients {
		ingredient, _ := strconv.Atoi(ingredient)
		fresh := false

		for _, freshRange := range freshRanges {
			if inRange(ingredient, freshRange) {
				fresh = true
				break
			}
		}

		if fresh {
			count++
		}
	}

	return count
}

func inRange(ingredient int, freshRange string) bool {
	parts := strings.Split(freshRange, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return ingredient >= start && ingredient <= end
}

func solvePart2(freshRanges []string) int {
	count := 0

	// sort ranges by start value
	sort.Slice(freshRanges, func(a, b int) bool {
		aParts := strings.Split(freshRanges[a], "-")
		bParts := strings.Split(freshRanges[b], "-")

		aStart, _ := strconv.Atoi(aParts[0])
		bStart, _ := strconv.Atoi(bParts[0])

		return aStart < bStart
	})

	// merge overlapping ranges
	currentRange := freshRanges[0]

	for i := 1; i < len(freshRanges); i++ {
		currentParts := strings.Split(currentRange, "-")
		currentStart, _ := strconv.Atoi(currentParts[0])
		currentEnd, _ := strconv.Atoi(currentParts[1])

		nextParts := strings.Split(freshRanges[i], "-")
		nextStart, _ := strconv.Atoi(nextParts[0])
		nextEnd, _ := strconv.Atoi(nextParts[1])

		if nextStart <= currentEnd {
			// ranges overlap, merge them
			if nextEnd > currentEnd {
				currentEnd = nextEnd
			}
			currentRange = strconv.Itoa(currentStart) + "-" + strconv.Itoa(currentEnd)
		} else {
			// no overlap, move to next range
			count += currentEnd - currentStart + 1

			currentRange = freshRanges[i]
		}
	}

	currentParts := strings.Split(currentRange, "-")
	currentStart, _ := strconv.Atoi(currentParts[0])
	currentEnd, _ := strconv.Atoi(currentParts[1])

	count += currentEnd - currentStart + 1

	return count
}
