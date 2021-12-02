package day2

import (
	util "aoc-2021/v2/utils"
	"log"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day2/input.txt")

func Part1() int {
	var depth int = 0
	var horizontal int = 0
	for _, line := range lines {
		slices := strings.SplitN(line, " ", 2)
		val, err := strconv.Atoi(slices[1])
		if err != nil {
			log.Fatal("Failed to convern string")
		}
		switch slices[0] {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return depth * horizontal
}

func Part2() int {
	var depth int = 0
	var horizontal int = 0
	var aim int = 0
	for _, line := range lines {
		slices := strings.SplitN(line, " ", 2)
		val, err := strconv.Atoi(slices[1])
		if err != nil {
			log.Fatal("Failed to convern string")
		}
		switch slices[0] {
		case "forward":
			horizontal += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return depth * horizontal
}
