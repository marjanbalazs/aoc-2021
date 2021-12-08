package day8

import (
	util "aoc-2021/v2/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day8/input.txt")

var knownLenghts map[int]int = make(map[int]int)

var originalCombinations map[int][]string = make(map[int][]string)

func init() {
	knownLenghts[2] = 1
	knownLenghts[4] = 4
	knownLenghts[3] = 7
	knownLenghts[7] = 8

	originalCombinations[1] = []string{"c", "f"}
	originalCombinations[7] = []string{"a", "c", "f"}
	originalCombinations[4] = []string{"b", "c", "d", "f"}

	originalCombinations[2] = []string{"a", "c", "d", "e", "g"}
	originalCombinations[3] = []string{"a", "c", "d", "f", "g"}
	originalCombinations[5] = []string{"a", "b", "d", "f", "g"}

	originalCombinations[0] = []string{"a", "b", "c", "e", "f", "g"}
	originalCombinations[6] = []string{"a", "b", "d", "e", "f", "g"}
	originalCombinations[9] = []string{"a", "b", "c", "d", "f", "g"}

	originalCombinations[8] = []string{"a", "b", "c", "d", "e", "f", "g"}
}

func Part1() int {
	var inputs []string = make([]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.SplitAfter(line, "|")
		trimmed := strings.Trim(parts[1], " ")
		inputs = append(inputs, trimmed)
	}
	var cntr int = 0
	for _, input := range inputs {
		parts := strings.Split(input, " ")
		for _, p := range parts {
			if knownLenghts[len(p)] != 0 {
				cntr++
			}
		}
	}
	return cntr
}

func findRuneSubset(strarray []string, substr string) string {
	for _, elem := range strarray {
		var passed bool = true
		for _, r := range substr {
			if strings.Index(elem, string(r)) == -1 {
				passed = false
			}
		}
		if passed {
			return elem
		}
	}
	return ""
}

func filterString(strarray []string, s string) []string {
	var filtered []string = []string{}
	for _, elem := range strarray {
		if elem != s {
			filtered = append(filtered, elem)
		}
	}
	return filtered

}

func Part2() int {
	var codings [][]string = make([][]string, 0, len(lines))
	var outputs [][]string = make([][]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.SplitN(line, "|", 2)
		trimmedCoding := strings.Trim(parts[0], " ")
		trimmedOutput := strings.Trim(parts[1], " ")
		var codingArray []string = []string{}
		for _, coding := range strings.Split(trimmedCoding, " ") {
			codingArray = append(codingArray, coding)
		}
		codings = append(codings, codingArray)
		var outputArray []string = []string{}
		for _, output := range strings.Split(trimmedOutput, " ") {
			outputArray = append(outputArray, output)
		}
		outputs = append(outputs, outputArray)
	}
	var sum int = 0
	for i := 0; i < len(lines); i++ {
		var lenmap map[int][]string = make(map[int][]string)
		for _, elem := range codings[i] {
			elemLen := len(elem)
			lenmap[elemLen] = append(lenmap[elemLen], elem)
		}
		var decoded map[string]string = make(map[string]string)

		decoded[lenmap[2][0]] = "1"
		decoded[lenmap[3][0]] = "7"
		decoded[lenmap[7][0]] = "8"
		decoded[lenmap[4][0]] = "4"

		// 6 length items
		nine := findRuneSubset(lenmap[6], lenmap[4][0])
		decoded[nine] = "9"
		lenmap[6] = filterString(lenmap[6], nine)

		zero := findRuneSubset(lenmap[6], lenmap[3][0])
		decoded[zero] = "0"
		lenmap[6] = filterString(lenmap[6], zero)

		decoded[lenmap[6][0]] = "6"

		// 5 length items
		three := findRuneSubset(lenmap[5], lenmap[3][0])
		decoded[three] = "3"
		lenmap[5] = filterString(lenmap[5], three)

		fiveOrNil := findRuneSubset(lenmap[6], lenmap[5][0])
		if fiveOrNil != "" {
			decoded[lenmap[5][0]] = "5"
			decoded[lenmap[5][1]] = "2"
		} else {
			decoded[lenmap[5][0]] = "2"
			decoded[lenmap[5][1]] = "5"
		}
		var outputVal string = ""
		for _, outputElem := range outputs[i] {
			for key, value := range decoded {
				splitKey := strings.Split(key, "")
				splitElem := strings.Split(outputElem, "")
				sort.Strings(splitKey)
				sort.Strings(splitElem)
				if strings.Join(splitKey, "") == strings.Join(splitElem, "") {
					outputVal += value
				}
			}
		}
		val, err := strconv.Atoi(outputVal)
		if err != nil {
			log.Fatal("Failed to parse int")
		}
		sum += val
	}
	return sum
}
