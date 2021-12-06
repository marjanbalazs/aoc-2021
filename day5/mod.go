package day5

import (
	util "aoc-2021/v2/utils"
	"log"
	"strconv"
	"strings"
)

var vents = util.ReadFile("./day5/input.txt")

type Point struct {
	x, y int
}

type Line struct {
	start Point
	end   Point
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (line Line) isVertical() bool {
	return line.start.x == line.end.x
}

func (line Line) isHorizontal() bool {
	return line.start.y == line.end.y
}

func VectorAdd(v1 Point, v2 Point) Point {
	return Point{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func (line Line) coveredPoints() []Point {
	points := []Point{}
	if line.isHorizontal() {
		length := line.end.x - line.start.x
		var v2 Point
		if length > 0 {
			v2 = Point{1, 0}
		} else {
			v2 = Point{-1, 0}
		}
		points = append(points, line.start)
		for i := 1; i < abs(length); i++ {
			lastPoint := VectorAdd(points[len(points)-1], v2)
			points = append(points, lastPoint)
		}
		points = append(points, line.end)
	} else if line.isVertical() {
		length := line.end.y - line.start.y
		var v2 Point
		if length > 0 {
			v2 = Point{0, 1}
		} else {
			v2 = Point{0, -1}
		}
		points = append(points, line.start)
		for i := 1; i < abs(length); i++ {
			lastPoint := VectorAdd(points[len(points)-1], v2)
			points = append(points, lastPoint)
		}
		points = append(points, line.end)
	} else {
		//This must be diagonal
		xdir := line.end.x - line.start.x
		ydir := line.end.y - line.start.y
		p := Point{}
		if xdir > 0 {
			p.x = 1
		} else {
			p.x = -1
		}
		if ydir > 0 {
			p.y = 1
		} else {
			p.y = -1
		}
		points = append(points, line.start)
		for ok := true; ok; ok = !(points[len(points)-1].x == line.end.x && points[len(points)-1].y == line.end.y) {
			lastPoint := VectorAdd(points[len(points)-1], p)
			points = append(points, lastPoint)
		}
	}
	return points
}

func parseLine(line string) Line {
	fragments := strings.Split(line, " -> ")
	start := strings.Split(fragments[0], ",")
	start_x, err := strconv.Atoi(start[0])
	if err != nil {
		log.Fatal("Failed to convert string to number")
	}
	start_y, err := strconv.Atoi(start[1])
	if err != nil {
		log.Fatal("Failed to convert string to number")
	}
	end := strings.Split(fragments[1], ",")
	end_x, err := strconv.Atoi(end[0])
	if err != nil {
		log.Fatal("Failed to convert string to number")
	}
	end_y, err := strconv.Atoi(end[1])
	if err != nil {
		log.Fatal("Failed to convert string to number")
	}
	return Line{
		start: Point{x: start_x, y: start_y},
		end:   Point{x: end_x, y: end_y},
	}
}

func Part1() int {
	lines := []Line{}
	for _, line := range vents {
		lines = append(lines, parseLine(line))
	}
	horizontalOrVertical := []Line{}
	for _, line := range lines {
		if line.isHorizontal() || line.isVertical() {
			horizontalOrVertical = append(horizontalOrVertical, line)
		}
	}
	covered := []Point{}
	for _, hvLine := range horizontalOrVertical {
		for _, coveredPoint := range hvLine.coveredPoints() {
			covered = append(covered, coveredPoint)
		}
	}
	cntr := map[string]int{}
	for _, c := range covered {
		x := strconv.Itoa(c.x)
		y := strconv.Itoa(c.y)
		key := strings.Join([]string{x, y}, ",")
		cntr[key] += 1
	}
	var ret int
	for _, value := range cntr {
		if value > 1 {
			ret++
		}
	}
	return ret
}

func Part2() int {
	lines := []Line{}
	for _, line := range vents {
		lines = append(lines, parseLine(line))
	}
	covered := []Point{}
	for _, line := range lines {
		for _, coveredPoint := range line.coveredPoints() {
			covered = append(covered, coveredPoint)
		}
	}
	cntr := map[string]int{}
	for _, c := range covered {
		x := strconv.Itoa(c.x)
		y := strconv.Itoa(c.y)
		key := strings.Join([]string{x, y}, ",")
		cntr[key] += 1
	}
	var ret int
	for _, value := range cntr {
		if value > 1 {
			ret++
		}
	}
	return ret
}
