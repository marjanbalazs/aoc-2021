package day10

import (
	util "aoc-2021/v2/utils"
	"errors"
	"sort"
	"strings"
)

var pointmap map[string]int = make(map[string]int)
var secondpointmap map[string]int = make(map[string]int)
var leftmap map[string]string = make(map[string]string)

func init() {
	pointmap[")"] = 3
	pointmap["]"] = 57
	pointmap["}"] = 1197
	pointmap[">"] = 25137

	leftmap["("] = ")"
	leftmap["["] = "]"
	leftmap["{"] = "}"
	leftmap["<"] = ">"

	secondpointmap["("] = 1
	secondpointmap["["] = 2
	secondpointmap["{"] = 3
	secondpointmap["<"] = 4
}

var lines = util.ReadFile("./day10/input.txt")

type Stack struct {
	arr []string
}

func (s *Stack) Pop() (string, error) {
	arrLen := len(s.arr)
	if arrLen > 0 {
		pop := s.arr[arrLen-1]
		s.arr = s.arr[:arrLen-1]
		return pop, nil
	}
	return "", errors.New("No elems to pop")
}

func (s *Stack) Push(elem string) {
	s.arr = append(s.arr, elem)
}

func (s *Stack) Peek() string {
	if len(s.arr) == 0 {
		return ""
	}
	return s.arr[len(s.arr)-1]
}

func Part1() int {
	var point int = 0
	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")
		stack := Stack{make([]string, 0)}
		for j := 0; j < len(chars); j++ {
			c := chars[j]
			if leftmap[c] != "" {
				stack.Push(c)
			} else {
				v := stack.Peek()
				if leftmap[v] == c {
					stack.Pop()
				} else {
					point += pointmap[c]
					break
				}
			}
		}
	}
	return point
}

func Part2() int {
	totalscores := make([]int, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")
		stack := Stack{make([]string, 0)}
		linescore := 0
		for j := 0; j < len(chars); j++ {
			c := chars[j]
			if leftmap[c] != "" {
				stack.Push(c)
			} else {
				v := stack.Peek()
				if leftmap[v] == c {
					stack.Pop()
				} else {
					break
				}
			}
			if j == len(chars)-1 {
				for stack.Peek() != "" {
					popped, _ := stack.Pop()
					value := secondpointmap[popped]
					linescore = linescore*5 + value
				}
			}
		}
		totalscores = append(totalscores, linescore)
	}
	filteredscores := make([]int, 0)
	for _, elem := range totalscores {
		if elem != 0 {
			filteredscores = append(filteredscores, elem)
		}
	}
	sort.Ints(filteredscores)
	return filteredscores[len(filteredscores)/2]
}
