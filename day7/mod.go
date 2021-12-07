package day7

import (
	util "aoc-2021/v2/utils"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

var line = util.ReadFile("./day7/input.txt")

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part1() int {
	stringPositions := strings.Split(line[0], ",")
	intPositions := make([]int, 0, len(stringPositions))
	for _, s := range stringPositions {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Failed to convert string to number")
		}
		intPositions = append(intPositions, val)
	}
	sort.Ints(intPositions)
	var sum int
	for _, val := range intPositions {
		sum += val
	}
	min := intPositions[0]
	max := intPositions[len(intPositions)-1]
	currentLowest := math.MaxInt64
	for i := min; i != max; i++ {
		var fuelSum int
		for _, val := range intPositions {
			fuelSum += abs(i - val)
		}
		if fuelSum < currentLowest {
			currentLowest = fuelSum
		}
	}
	return currentLowest
}

func fuelCostFunction(len int) int {
	var sum int = 0
	for i := 0; i <= len; i++ {
		sum += i
	}
	return sum
}

func Part2() int {
	stringPositions := strings.Split(line[0], ",")
	intPositions := make([]int, 0, len(stringPositions))
	for _, s := range stringPositions {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Failed to convert string to number")
		}
		intPositions = append(intPositions, val)
	}
	sort.Ints(intPositions)
	var sum int
	for _, val := range intPositions {
		sum += val
	}
	min := intPositions[0]
	max := intPositions[len(intPositions)-1]
	currentLowest := math.MaxInt64
	for i := min; i <= max; i++ {
		var fuelSum int
		for _, val := range intPositions {
			fuelSum += fuelCostFunction(abs(i - val))
		}
		if fuelSum < currentLowest {
			currentLowest = fuelSum
		}
	}
	return currentLowest
}
