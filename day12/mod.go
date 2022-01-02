package day12

import (
	util "aoc-2021/v2/utils"
	"strings"
	"unicode"
)

var lines = util.ReadFile("./day12/input.txt")

type Cave struct {
	neighbours map[string]bool
}

type CaveStructure map[string]Cave

func findOccurances(s string, arr []string) int {
	cntr := 0
	for _, elem := range arr {
		if elem == s {
			cntr += 1
		}
	}
	return cntr
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func hasString(s string, arr []string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func addCavePair(c1 string, c2 string, caves *CaveStructure) {
	cavestruct := *caves
	_, contined := cavestruct[c1]
	if contined {
		cave := cavestruct[c1]
		cave.neighbours[c2] = true
	} else {
		neighbours := make(map[string]bool)
		neighbours[c2] = true
		cave := Cave{neighbours: neighbours}
		cavestruct[c1] = cave
	}
}

func buildCaveStructure() CaveStructure {
	caveStructure := make(CaveStructure)
	for _, line := range lines {
		fragments := strings.SplitN(line, "-", 2)
		addCavePair(fragments[0], fragments[1], &caveStructure)
		addCavePair(fragments[1], fragments[0], &caveStructure)
	}
	return caveStructure
}

func seekEnd(currCave string, caves CaveStructure, path []string, endPaths [][]string) [][]string {
	for key := range caves[currCave].neighbours {
		switch key {
		case "end":
			{
				endPath := []string{}
				for _, v := range path {
					endPath = append(endPath, v)
				}
				endPath = append(endPath, key)
				endPaths = append(endPaths, endPath)
				break
			}
		case "start":
			{
				break
			}
		default:
			{
				blocked := hasString(key, path) && !isUpper(key)
				if !blocked {
					var newPath []string = []string{}
					for _, v := range path {
						newPath = append(newPath, v)
					}
					newPath = append(newPath, key)
					endPaths = seekEnd(key, caves, newPath, endPaths)
				}
			}
		}
	}
	return endPaths
}

// This could be made prettier by passing a conditiong checker function
func seekEndWithDuplicate(currCave string, duplicate string, caves CaveStructure, path []string, endPaths [][]string) [][]string {
	for key := range caves[currCave].neighbours {
		switch key {
		case "end":
			{
				endPath := []string{}
				for _, v := range path {
					endPath = append(endPath, v)
				}
				endPath = append(endPath, key)
				endPaths = append(endPaths, endPath)
				break
			}
		case "start":
			{
				break
			}
		default:
			{
				occurances := findOccurances(key, path)
				secondChance := occurances == 1 && duplicate == key
				if isUpper(key) || occurances == 0 || secondChance {
					var newPath []string = []string{}
					for _, v := range path {
						newPath = append(newPath, v)
					}
					newPath = append(newPath, key)
					endPaths = seekEndWithDuplicate(key, duplicate, caves, newPath, endPaths)
				}
			}
		}
	}
	return endPaths
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Part1() int {
	caves := buildCaveStructure()
	path := []string{"start"}
	endPaths := make([][]string, 0)
	endPaths = seekEnd("start", caves, path, endPaths)
	return len(endPaths)
}

func Part2() int {
	caves := buildCaveStructure()
	results := []string{}
	// A lot of time could be saved by checking for duplicates along the path inside the seeking funcion
	for caveName, _ := range caves {
		if caveName != "start" && caveName != "end" && !isUpper(caveName) {
			path := []string{"start"}
			endPaths := make([][]string, 0)
			endPaths = seekEndWithDuplicate("start", caveName, caves, path, endPaths)
			concat := []string{}
			for _, arr := range endPaths {
				concat = append(concat, strings.Join(arr, ","))
			}
			results = append(results, concat...)
		}
	}
	duplicatesRemoved := removeDuplicateStr(results)
	return len(duplicatesRemoved)
}
