package day3

import (
	util "aoc-2021/v2/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day3/input.txt")

func countOnes(lines []string) []string {
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
		} else {
			mostcommonbits[i] = "0"
		}
	}
	return mostcommonbits
}

func Part1() int64 {
	gammabits := countOnes(lines)
	epsilonbits := make([]string, len(gammabits))
	for i, bit := range gammabits {
		if bit == "0" {
			epsilonbits[i] = "1"
		} else {
			epsilonbits[i] = "0"
		}
	}
	gamma, err := strconv.ParseInt(strings.Join(gammabits, ""), 2, 32)
	epsilon, err := strconv.ParseInt(strings.Join(epsilonbits, ""), 2, 32)
	if err != nil {
		log.Fatal("Failed to convert string")
	}
	return gamma * epsilon
}

func Part2() int {
	mostcommon := countOnes(lines)
	leastcommon := make([]string, len(mostcommon))
	for i, bit := range mostcommon {
		if bit == "0" {
			leastcommon[i] = "1"
		} else {
			leastcommon[i] = "0"
		}
	}

	oxygenMax, err := strconv.ParseInt(strings.Join(mostcommon, ""), 2, 31)
	co2Low, err := strconv.ParseInt(strings.Join(leastcommon, ""), 2, 31)

	if err != nil {
		log.Fatal("Failed to convert string")
	}

	fmt.Println(oxygenMax)
	fmt.Println(co2Low)
	var linevalues []int
	for _, line := range lines {
		val, err := strconv.ParseInt(line, 2, 31)
		if err != nil {
			log.Fatal("Failed to convert string")
		}
		linevalues = append(linevalues, int(val))
	}
	sort.Ints(linevalues)
	var oxygen int
	var co2 int
	fmt.Println("Sorted")
	fmt.Println(linevalues[0])
	fmt.Println(linevalues[999])
	for i := 1; i < len(linevalues); i++ {
		if linevalues[i] > int(co2Low) {
			co2 = linevalues[i-1]
			break
		}
	}

	for _, val := range linevalues {
		if val > int(oxygenMax) {
			oxygen = val
			break
		}
	}

	fmt.Println(co2, oxygen)
	return oxygen * co2
}
