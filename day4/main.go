package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"pcancoin/hello/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	total := 0
	for _, shifts := range parseInput(input) {
		total += checkOverlap1(shifts)
	}
	return total
}

func part2(input string) int {
	total := 0
	for _, shifts := range parseInput(input) {
		total += checkOverlap2(shifts)
	}
	return total
}

func parseInput(input string) [][][]int {
	parsedInput := make([][][]int, 0)
	for _, l := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, [][]int{})
		for _, shift := range strings.Split(l, ",") {
			times := strings.Split(shift, "-")
			time1, _ := strconv.Atoi(times[0])
			time2, _ := strconv.Atoi(times[1])
			parsedInput[len(parsedInput)-1] = append(parsedInput[len(parsedInput)-1], []int{time1, time2})
		}
	}
	return parsedInput
}

func checkOverlap1(shifts [][]int) int {
	if (shifts[0][0] >= shifts[1][0] && shifts[0][1] <= shifts[1][1]) ||
		(shifts[1][0] >= shifts[0][0] && shifts[1][1] <= shifts[0][1]) {
		return 1
	}
	return 0
}

func checkOverlap2(shifts [][]int) int {
	if (shifts[0][0] >= shifts[1][0] && shifts[0][0] <= shifts[1][1]) ||
		(shifts[0][1] >= shifts[1][0] && shifts[0][1] <= shifts[1][1]) ||
		(shifts[0][0] <= shifts[1][0] && shifts[0][1] >= shifts[1][1]) {
		return 1
	}
	return 0
}
