package day4

import (
	util "aoc-2021/v2/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadFile("./day4/input.txt")

type board [][]string

func generateBoards(lines []string) []board {
	boards := []board{}

	i := 0
	for i < len(lines) {
		if len(lines[i]) == 0 {
			i++
			continue
		}
		board := board{}
		for i < len(lines) && len(lines[i]) != 0 {
			values := []string{}
			for _, fragment := range strings.Split(lines[i], " ") {
				if len(fragment) != 0 {
					values = append(values, strings.TrimSpace(fragment))
				}
			}
			board = append(board, values)
			i++
		}
		boards = append(boards, board)
	}
	return boards
}

func isWinner(numbers []string, candidate []string) bool {
	set := map[string]bool{}
	for _, number := range numbers {
		set[number] = true
	}
	for _, elem := range candidate {
		if !set[elem] {
			return false
		}
	}
	return true
}

func checkRows(numbers []string, board board) bool {
	for _, row := range board {
		if isWinner(numbers, row) {
			return true
		}
	}
	return false
}

func checkColumns(numbers []string, board board) bool {
	for i := 0; i < len(board); i++ {
		column := []string{}
		for _, row := range board {
			column = append(column, row[i])
		}
		if isWinner(numbers, column) {
			return true
		}
	}
	return false
}

func getWinningCombination(numbers []string, boards []board) ([]string, board) {
	for i := 5; i < len(numbers); i++ {
		currentNumbers := numbers[:i]
		for _, board := range boards {
			if checkColumns(currentNumbers, board) || checkRows(currentNumbers, board) {
				return currentNumbers, board
			}
		}
	}
	return []string{"-1"}, boards[0]
}

func boardValue(sequence []string, board board) int {
	sum := 0
	winningSet := map[string]bool{}
	for _, number := range sequence {
		winningSet[number] = true
	}
	for _, row := range board {
		for _, elem := range row {
			if !winningSet[string(elem)] {
				number, err := strconv.Atoi(strings.TrimSpace(elem))
				if err != nil {
					log.Fatal(err)
				}
				sum += number
			}
		}
	}

	lastValue, err := strconv.Atoi(sequence[len(sequence)-1])
	if err != nil {
		log.Fatal("Failed to convert string to int")
	}
	return sum * lastValue
}

func Part1() int {
	numbers := strings.Split(lines[0], ",")
	boards := generateBoards(lines[2:])
	winningSequence, winningBoard := getWinningCombination(numbers, boards)
	return boardValue(winningSequence, winningBoard)
}

type Winner struct {
	sequence []string
	board    board
}

func Part2() int {
	numbers := strings.Split(lines[0], ",")
	boards := generateBoards(lines[2:])
	winners := []Winner{}
	for _, board_elem := range boards {
		thisIteration := []board{board_elem}
		winningSequence, winningBoard := getWinningCombination(numbers, thisIteration)
		winners = append(winners, Winner{winningSequence, winningBoard})
	}

	sort.SliceStable(winners, func(i, j int) bool {
		return len(winners[i].sequence) < len(winners[j].sequence)
	})

	lastWinnerSequence := winners[len(winners)-1].sequence
	lastWinnerBoard := winners[len(winners)-1].board
	return boardValue(lastWinnerSequence, lastWinnerBoard)
}
