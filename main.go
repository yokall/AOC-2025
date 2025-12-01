package main

import (
	"log"
	"path/filepath"

	. "github.com/yokall/aoc-2025/day01"
	. "github.com/yokall/aoc-2025/utils"
)

func main() {
	inputPath, err := filepath.Abs("inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	input, err := ReadRawFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	part1, part2 := Solve(input)

	log.Printf("Day 01 - Part 1: %d", part1)
	log.Printf("Day 01 - Part 2: %d", part2)
}
