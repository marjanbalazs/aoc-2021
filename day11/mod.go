package day11

import (
	util "aoc-2021/v2/utils"
	"log"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day11/input.txt")

type Coordinate struct {
	x, y int
}

type Map struct {
	points [][]int
}

func NewMap(p [][]int) *Map {
	m := new(Map)
	m.points = p
	return m
}

func (m *Map) getHeight() int {
	return len(m.points)
}

func (m *Map) getWidht() int {
	return len(m.points[0])
}

func (m *Map) getValue(x int, y int) int {
	return m.points[y][x]
}

func (m *Map) setValue(x int, y int, val int) {
	m.points[y][x] = val
}

func (m *Map) getNeighbours(x int, y int) []Coordinate {
	height := len(m.points)
	widht := len(m.points[0])
	possibleNeighbours := []Coordinate{
		{x - 1, y + 1},
		{x - 1, y},
		{x - 1, y - 1},
		{x, y + 1},
		{x, y - 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x + 1, y - 1},
	}
	neighbours := []Coordinate{}
	for _, p := range possibleNeighbours {
		if p.x >= 0 && p.x < widht && p.y >= 0 && p.y < height {
			neighbours = append(neighbours, p)
		}
	}
	return neighbours
}

func flashPoint(x int, y int, m *Map) {
	m.setValue(x, y, 0)
	neighbours := m.getNeighbours(x, y)
	for _, n := range neighbours {
		c := m.getValue(n.x, n.y)
		if c != 0 {
			m.setValue(n.x, n.y, c+1)
		}
	}
}

func Part1() int {
	var octopuses [][]int = make([][]int, 0, len(lines))
	for i, line := range lines {
		oc := strings.Split(line, "")
		octopuses = append(octopuses, make([]int, 0))
		for _, o := range oc {
			val, err := strconv.Atoi(o)
			if err != nil {
				log.Fatal("Failed to conver string to int")
			}
			octopuses[i] = append(octopuses[i], val)
		}
	}
	flashes := 0
	m := NewMap(octopuses)
	for i := 0; i < 100; i++ {
		// Increment every point
		for k, line := range m.points {
			for j, _ := range line {
				val := m.getValue(j, k)
				m.setValue(j, k, val+1)
			}
		}
		// Flash once
		for y, l := range m.points {
			for x := range l {
				value := m.getValue(x, y)
				if value > 9 {
					flashPoint(x, y, m)
					flashes++
				}
			}
		}
		// Flash until everyone flashed
		for true {
			currentFlashes := flashes
			for y, l := range m.points {
				for x := range l {
					value := m.getValue(x, y)
					if value > 9 {
						flashPoint(x, y, m)
						flashes++
					}
				}
			}
			if currentFlashes == flashes {
				break
			}
		}

	}
	return flashes
}

func Part2() int {
	var octopuses [][]int = make([][]int, 0, len(lines))
	for i, line := range lines {
		oc := strings.Split(line, "")
		octopuses = append(octopuses, make([]int, 0))
		for _, o := range oc {
			val, err := strconv.Atoi(o)
			if err != nil {
				log.Fatal("Failed to conver string to int")
			}
			octopuses[i] = append(octopuses[i], val)
		}
	}
	flashes := 0
	m := NewMap(octopuses)
	allOctopuses := m.getHeight() * m.getWidht()
	for i := 0; i < 1000; i++ {
		originalFlashes := flashes
		// Increment every point
		for k, line := range m.points {
			for j, _ := range line {
				val := m.getValue(j, k)
				m.setValue(j, k, val+1)
			}
		}
		// Flash once
		for y, l := range m.points {
			for x := range l {
				value := m.getValue(x, y)
				if value > 9 {
					flashPoint(x, y, m)
					flashes++
				}
			}
		}
		// Flash until everyone flashed
		for true {
			currentFlashes := flashes
			for y, l := range m.points {
				for x := range l {
					value := m.getValue(x, y)
					if value > 9 {
						flashPoint(x, y, m)
						flashes++
					}
				}
			}
			if currentFlashes == flashes {
				break
			}
		}
		if originalFlashes+allOctopuses == flashes {
			return i + 1
		}
	}
	return flashes
}
