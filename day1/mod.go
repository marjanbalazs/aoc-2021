package day1

import (
	util "aoc-2021/v2/utils"
	"fmt"
	"log"
	"strconv"
)

func Part1() int {
	lines := util.ReadFile("./day1/input.txt")
	var acc int = 0
	for i := 1; i < len(lines); i++ {
		prev, err := strconv.Atoi(lines[i-1])
		curr, err := strconv.Atoi(lines[i])
		if err != nil {
			fmt.Println(err)
			log.Fatal("Failed to convern the string")
		}
		if prev < curr {
			acc++
		}
	}
	return acc
}

func Part2() int {
	lines := util.ReadFile("./day1/input.txt")
	var ints []int
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Failed to convert the string")
		}
		ints = append(ints, val)
	}

	var acc int = 0
	for i := 3; i < len(ints); i++ {
		prev := ints[i-3] + ints[i-2] + ints[i-1]
		curr := ints[i-2] + ints[i-1] + ints[i-0]
		if prev < curr {
			acc++
		}
	}
	return acc
}
