package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"pcancoin/hello/util"
)

type Move struct {
	itemCount int
	from      int
	to        int
}

type Stack []string
type Stacks []Stack

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

func part1(input string) string {
	stacks, moves := parseInput(input)

	for _, move := range moves {
		applyMove9000(stacks, move)
	}

	fmt.Println(stacks)

	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}

func part2(input string) string {
	stacks, moves := parseInput(input)

	for _, move := range moves {
		applyMove9001(stacks, move)
	}

	fmt.Println(stacks)

	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}

func parseInput(input string) (stacks Stacks, moves []Move) {
	// parsedInput := make([][]int, 0)

	inputAndMove := strings.Split(input, "\n\n")

	stacks = parseStacks(inputAndMove[0])
	moves = parseMoves(inputAndMove[1])

	return
}

func parseStacks(stacksInput string) (stacks Stacks) {
	input := strings.Split(stacksInput, "\n")

	tmp := strings.Split(strings.Trim(input[len(input)-1], " "), "   ")
	stacksCount, _ := strconv.Atoi(tmp[len(tmp)-1])

	stacks = Stacks{}
	for i := 0; i < stacksCount; i++ {
		stacks = append(stacks, Stack{})
	}

	for lineIndex := len(input) - 2; lineIndex >= 0; lineIndex-- {
		line := input[lineIndex]
		for stackIndex := 0; stackIndex < stacksCount; stackIndex++ {
			item := strings.Split(line[stackIndex*4:stackIndex*4+3], "")
			if item[0] == "[" {
				stacks[stackIndex] = append(stacks[stackIndex], item[1])
			}
		}
	}

	return
}

func parseMoves(movesInput string) (moves []Move) {
	input := strings.Split(movesInput, "\n")

	moves = []Move{}

	for _, moveLine := range input {
		move := strings.Split(moveLine, " ")

		itemCount, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		to, _ := strconv.Atoi(move[5])

		moves = append(moves, Move{
			itemCount,
			from - 1,
			to - 1,
		})
	}

	return
}

func applyMove9000(stacks Stacks, move Move) {
	for i := 0; i < move.itemCount; i++ {
		stacks[move.to] = append(stacks[move.to], stacks[move.from][len(stacks[move.from])-1])
		stacks[move.from] = stacks[move.from][:len(stacks[move.from])-1]
	}
}

func applyMove9001(stacks Stacks, move Move) {
	for _, item := range stacks[move.from][len(stacks[move.from])-move.itemCount:] {
		stacks[move.to] = append(stacks[move.to], item)
	}
	stacks[move.from] = stacks[move.from][:len(stacks[move.from])-move.itemCount]
}
