package day6

import (
	util "aoc-2021/v2/utils"
	"log"
	"strconv"
	"strings"
)

var line = util.ReadFile("./day6/input.txt")

type LanternFish struct {
	timer int
}

func Part1() int {
	intervals := strings.Split(line[0], ",")
	lanternfish := []LanternFish{}
	for _, interval := range intervals {
		val, err := strconv.Atoi(interval)
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		lanternfish = append(lanternfish, LanternFish{val})
	}
	for i := 0; i < 80; i++ {
		babies := []LanternFish{}
		for i, _ := range lanternfish {
			if lanternfish[i].timer == 0 {
				babies = append(babies, LanternFish{8})
				lanternfish[i].timer = 6
			} else {
				lanternfish[i].timer = lanternfish[i].timer - 1
			}
		}
		lanternfish = append(lanternfish, babies...)

	}
	return len(lanternfish)
}

func Part2() int {
	intervals := strings.Split(line[0], ",")
	lanternfish := []int{}
	for _, interval := range intervals {
		val, err := strconv.Atoi(interval)
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		lanternfish = append(lanternfish, val)
	}
	fishIntervalArray := make([]int, 9)
	for _, fish := range lanternfish {
		fishIntervalArray[fish]++
	}
	for i := 0; i < 256; i++ {
		newBred := fishIntervalArray[0]
		fishIntervalArray[0] = fishIntervalArray[1]
		fishIntervalArray[1] = fishIntervalArray[2]
		fishIntervalArray[2] = fishIntervalArray[3]
		fishIntervalArray[3] = fishIntervalArray[4]
		fishIntervalArray[4] = fishIntervalArray[5]
		fishIntervalArray[5] = fishIntervalArray[6]
		fishIntervalArray[6] = fishIntervalArray[7] + newBred
		fishIntervalArray[7] = fishIntervalArray[8]
		fishIntervalArray[8] = newBred
	}
	var sum int
	for _, val := range fishIntervalArray {
		sum += val
	}
	return sum
}
