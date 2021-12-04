package day3

import (
	util "aoc-2021/v2/utils"
	"log"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day3/input.txt")

func countMostCommon(lines []string) []string {
	threshold := len(lines) / 2
	counters := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, ch := range line {
			if string(ch) == "1" {
				counters[i]++
			}
		}
	}
	mostcommonbits := make([]string, len(counters))
	for i, cntr := range counters {
		if cntr > threshold {
			mostcommonbits[i] = "1"
		} else if cntr == threshold {
			mostcommonbits[i] = "X"
		} else {
			mostcommonbits[i] = "0"
		}
	}
	return mostcommonbits
}

func countLeastCommon(lines []string) []string {
	mostcommon := countMostCommon(lines)
	leastcommon := []string{}
	for _, val := range mostcommon {
		if val == "X" {
			leastcommon = append(leastcommon, "X")
		}
		if val == "1" {
			leastcommon = append(leastcommon, "0")
		}
		if val == "0" {
			leastcommon = append(leastcommon, "1")
		}
	}
	return leastcommon
}

func Part1() int64 {
	gammabits := countMostCommon(lines)
	epsilonbits := countLeastCommon(lines)
	gamma, err := strconv.ParseInt(strings.Join(gammabits, ""), 2, 32)
	epsilon, err := strconv.ParseInt(strings.Join(epsilonbits, ""), 2, 32)
	if err != nil {
		log.Fatal("Failed to convert string")
	}
	return gamma * epsilon
}

func filterForNumber(currentBit string, currentPos int, lines []string, def string) []string {
	filtered := []string{}
	for _, line := range lines {
		if currentBit == "X" {
			if string([]rune(line)[currentPos]) == def {
				filtered = append(filtered, line)
			}
		} else {
			if string([]rune(line)[currentPos]) == currentBit {
				filtered = append(filtered, line)
			}
		}
	}
	return filtered
}

func Part2() int {
	var oxygenFiltered = lines
	i := 0
	for len(oxygenFiltered) != 1 {
		mostcommonbits := countMostCommon(oxygenFiltered)
		oxygenFiltered = filterForNumber(mostcommonbits[i], i, oxygenFiltered, "1")
		i++
	}
	o2, err := strconv.ParseInt(oxygenFiltered[0], 2, 15)
	if err != nil {
		log.Fatal("Failed to convert")
	}
	var co2Filtered = lines
	j := 0
	for len(co2Filtered) != 1 {
		leastcommon := countLeastCommon(co2Filtered)
		co2Filtered = filterForNumber(leastcommon[j], j, co2Filtered, "0")
		j++
	}
	co2, err := strconv.ParseInt(co2Filtered[0], 2, 15)
	if err != nil {
		log.Fatal("Failed to convert")
	}
	return int(co2 * o2)
}
