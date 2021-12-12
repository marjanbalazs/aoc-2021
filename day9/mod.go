package day9

import (
	util "aoc-2021/v2/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day9/input.txt")

type Point struct {
	x, y int
}

func getNeighbours(x int, y int, height int, widht int) []Point {
	var neighbours = make([]Point, 0, 4)
	if x == 0 {
		neighbours = append(neighbours, Point{x + 1, y})
	} else if x == widht-1 {
		neighbours = append(neighbours, Point{x - 1, y})
	} else {
		neighbours = append(neighbours, Point{x + 1, y})
		neighbours = append(neighbours, Point{x - 1, y})
	}

	if y == 0 {
		neighbours = append(neighbours, Point{x, y + 1})
	} else if y == height-1 {
		neighbours = append(neighbours, Point{x, y - 1})
	} else {
		neighbours = append(neighbours, Point{x, y + 1})
		neighbours = append(neighbours, Point{x, y - 1})
	}
	return neighbours
}

func isLowPoint(heightmap [][]int, x int, y int) (bool, int) {
	height := len(heightmap)
	widht := len(heightmap[0])
	neighbours := getNeighbours(x, y, height, widht)

	this := heightmap[y][x]
	var localMinima bool = true
	for _, n := range neighbours {
		if heightmap[n.y][n.x] <= this {
			localMinima = false
		}
	}
	if localMinima {
		return true, this + 1
	} else {
		return false, this + 1
	}
}

func Part1() int {
	heightmap := make([][]int, 0, len(lines))
	for i, line := range lines {
		values := strings.Split(line, "")
		heightmap = append(heightmap, make([]int, 0, len(values)))
		for _, value := range values {
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("Failed to convert string")
			}
			heightmap[i] = append(heightmap[i], val)
		}
	}

	sum := 0
	lowpoints := 0
	for y, l := range heightmap {
		for x := range l {
			lowpoint, risk := isLowPoint(heightmap, x, y)
			if lowpoint {
				lowpoints++
				sum += risk
			}
		}
	}
	return sum
}

func Part2() int {
	heightmap := make([][]int, 0, len(lines))
	for i, line := range lines {
		values := strings.Split(line, "")
		heightmap = append(heightmap, make([]int, 0, len(values)))
		for _, value := range values {
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("Failed to convert string")
			}
			heightmap[i] = append(heightmap[i], val)
		}
	}
	height := len(heightmap)
	widht := len(heightmap[0])
	var currentColor int = 10
	var cntr map[int]int = make(map[int]int)

	var lowpoints []Point = []Point{}
	for y, l := range heightmap {
		for x := range l {
			lowpoint, _ := isLowPoint(heightmap, x, y)
			if lowpoint {
				lowpoints = append(lowpoints, Point{x, y})
			}
		}
	}

	for _, lowpoint := range lowpoints {
		done := []Point{}
		toCheck := []Point{}
		toCheck = append(toCheck, lowpoint)
		for len(toCheck) > 0 {
			newPoints := []Point{}
			for _, c := range toCheck {
				neighbours := getNeighbours(c.x, c.y, height, widht)
				for _, n := range neighbours {
					if heightmap[n.y][n.x] < 9 {
						heightmap[n.y][n.x] = currentColor
						newPoints = append(newPoints, Point{n.x, n.y})
					}
				}
				done = append(done, c)
			}
			toCheck = newPoints
		}
		currentColor++
	}

	for _, l := range heightmap {
		for _, v := range l {
			if cntr[v] == 0 {
				cntr[v] = 1
			} else {
				cntr[v] += 1
			}
		}
	}

	var values []int = make([]int, 0, len(cntr))
	for key, value := range cntr {
		if key != 9 {
			values = append(values, value)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values[0] * values[1] * values[2]
}
