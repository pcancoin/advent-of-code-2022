package main

import (
	"flag"
	"fmt"
	"strings"

	"pcancoin/hello/cast"
	"pcancoin/hello/util"

	"github.com/juliangruber/go-intersect"
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

// ASCIICodeCapA   = int('A') // 65
// ASCIICodeCapZ   = int('Z') // 90
// ASCIICodeLowerA = int('a') // 97
// ASCIICodeLowerZ = int('z') // 122

func part1(input string) int {
	parsedInput := parseInput(input)
	total := 0
	for _, backpack := range parsedInput {
		compartment1 := strings.Split(backpack[:(len(backpack)/2)], "")
		compartment2 := strings.Split(backpack[(len(backpack)/2):], "")
		intersection := intersect.Hash(compartment1, compartment2)
		total += getPriority(fmt.Sprintf("%v", intersection[0]))
	}
	fmt.Println(cast.ASCIICodeCapZ)
	return total
}

func part2(input string) int {
	parsedInput := parseInput(input)
	total := 0
	for i := 0; i < len(parsedInput); i += 3 {
		backpack1 := strings.Split(parsedInput[i], "")
		backpack2 := strings.Split(parsedInput[i+1], "")
		backpack3 := strings.Split(parsedInput[i+2], "")
		intersection1 := intersect.Hash(backpack1, backpack2)
		intersection2 := intersect.Hash(intersection1, backpack3)
		total += getPriority(fmt.Sprintf("%v", intersection2[0]))
	}
	return total
}

func parseInput(input string) []string {
	parsedInput := make([]string, 0)
	for _, l := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, l)
	}
	return parsedInput
}

func getPriority(item string) int {
	asciicode := cast.ToASCIICode(item)
	if asciicode >= cast.ASCIICodeCapA && asciicode <= cast.ASCIICodeCapZ {
		return asciicode - cast.ASCIICodeCapA + 27
	}
	return asciicode - cast.ASCIICodeLowerA + 1
}
