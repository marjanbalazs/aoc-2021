package day12

import (
	util "aoc-2021/v2/utils"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

var lines = util.ReadFile("./day12/input.txt")

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

type Cave struct {
	neighbours map[string]bool
}

type CaveStructure map[string]Cave

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

func Part1() int {
	caves := buildCaveStructure()
	path := []string{"start"}
	endPaths := make([][]string, 0)
	endPaths = seekEnd("start", caves, path, endPaths)
	return len(endPaths)
}
