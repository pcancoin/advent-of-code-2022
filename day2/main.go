package main

import (
	"flag"
	"fmt"
	"strings"

	"pcancoin/hello/util"
)

type Sign int

const (
	Rock Sign = iota + 1
	Paper
	Scissors
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
	parsedInput := parseInput(input)
	total := 0
	for _, game := range parsedInput {
		total += computeScore1(game)
	}
	return total
}

func part2(input string) int {
	total := computeScore2(input)
	return total
}

func parseInput(input string) [][]Sign {
	parsedInput := make([][]Sign, 0)
	for _, l := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, make([]Sign, 0))
		for _, sign := range strings.Split(l, " ") {
			parsedInput[len(parsedInput)-1] = append(parsedInput[len(parsedInput)-1], charToSign(sign))
		}
	}
	return parsedInput
}

func computeScore2(input string) int {
	result := map[string]int{
		"A X": 3,
		"B X": 1,
		"C X": 2,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"A Z": 8,
		"B Z": 9,
		"C Z": 7,
	}
	total := 0
	for _, l := range strings.Split(input, "\n") {
		total += result[l]
	}
	return total
}

func charToSign(char string) Sign {
	switch char {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	fmt.Println(char)
	panic("Unknow entry : " + char)
}

func signToValue(sign Sign) int {
	switch sign {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic("Unknow sign")
}

func (e Sign) String() string {
	switch e {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func computeScore1(hand []Sign) int {
	if hand[1] == Rock && hand[0] == Scissors || hand[1] == Scissors && hand[0] == Paper || hand[1] == Paper && hand[0] == Rock {
		// fmt.Println(fmt.Sprintf("%s %s %d", hand[0].String(), hand[1].String(), 6+signToValue(hand[1])))
		return 6 + signToValue(hand[1])
	} else if hand[0] == hand[1] {
		// fmt.Println(fmt.Sprintf("%s %s %d", hand[0].String(), hand[1].String(), 3+signToValue(hand[1])))
		return 3 + signToValue(hand[1])
	} else {
		// fmt.Println(fmt.Sprintf("%s %s %d", hand[0].String(), hand[1].String(), signToValue(hand[1])))
		return signToValue(hand[1])
	}
}
