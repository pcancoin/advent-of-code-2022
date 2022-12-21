package main

import (
	"flag"
	"fmt"
	"strings"

	"pcancoin/hello/util"
)

func main() {
	var part int
	var useTestInput bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&useTestInput, "test", false, "Use input-test.txt")
	flag.Parse()
	fmt.Println("Running part", part)

	var input string

	if useTestInput {
		input = "./input-test.txt"
	} else {
		input = "./input.txt"
	}

	if part == 1 {
		ans := part1(util.ReadFile(input))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile(input))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	data := parseInput(input)

	fmt.Println(data)

	return 0
}

func part2(input string) int {
	data := parseInput(input)

	fmt.Println(data)

	return 0
}

func parseInput(input string) (data []string) {
	data = strings.Split(input, "\n")
	return
}
