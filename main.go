package main

import (
	"fmt"
	"log"
	"path/filepath"

	. "github.com/yokall/aoc-2025/day01"
	. "github.com/yokall/aoc-2025/day02"
	. "github.com/yokall/aoc-2025/day03"
	. "github.com/yokall/aoc-2025/day04"
	. "github.com/yokall/aoc-2025/day05"
	. "github.com/yokall/aoc-2025/utils"
)

var solvers = map[string]func(string) (int, int){
	"01": Solve,
	"02": SolveDay02,
	"03": SolveDay03,
	"04": SolveDay04,
	"05": SolveDay05,
}

func main() {
	// solveDay("01")
	// solveDay("02")
	// solveDay("03")
	// solveDay("04")
	solveDay("05")
}

func solveDay(day string) {
	inputPath, err := filepath.Abs(fmt.Sprintf("inputs/day%s.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	input, err := ReadRawFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	part1, part2 := solvers[day](input)

	log.Printf("Day %s - Part 1: %d", day, part1)
	log.Printf("Day %s - Part 2: %d", day, part2)
}
