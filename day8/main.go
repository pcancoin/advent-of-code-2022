package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"pcancoin/hello/util"
)

type Tree struct {
	height  int
	visible bool
}

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

	leftViewCompute(data)
	rightViewCompute(data)
	bottomViewCompute(data)
	topViewCompute(data)

	// printForest(data)

	return totalVisibleTrees(data)
}

func part2(input string) int {
	data := parseInput(input)

	fmt.Println(data)

	return 0
}

func parseInput(input string) (forest [][]*Tree) {
	parsed := strings.Split(input, "\n")

	forest = [][]*Tree{}

	for _, l := range parsed {
		forest = append(forest, []*Tree{})
		for _, t := range strings.Split(l, "") {
			height, _ := strconv.Atoi(t)
			forest[len(forest)-1] = append(forest[len(forest)-1], &Tree{
				height,
				false,
			})
		}
	}

	return
}

func leftViewCompute(forest [][]*Tree) [][]*Tree {
	for _, line := range forest {
		actualHeight := -1
		for _, tree := range line {
			if tree.height > actualHeight {
				actualHeight = tree.height
				tree.visible = true
			}
		}
	}
	return forest
}

func rightViewCompute(forest [][]*Tree) [][]*Tree {
	for _, line := range forest {
		actualHeight := -1

		for _, tree := range reverseLine(line) {
			if tree.height > actualHeight {
				actualHeight = tree.height
				tree.visible = true
			}
		}
	}
	return forest
}

func topViewCompute(forest [][]*Tree) [][]*Tree {
	for j := 0; j < len(forest[0]); j++ {
		actualHeight := -1
		for i := 0; i < len(forest); i++ {
			tree := forest[i][j]
			if tree.height > actualHeight {
				actualHeight = tree.height
				tree.visible = true
			}
		}
	}
	return forest
}

func bottomViewCompute(forest [][]*Tree) [][]*Tree {
	for j := 0; j < len(forest[0]); j++ {
		actualHeight := -1
		for i := len(forest) - 1; i > 0; i-- {
			tree := forest[i][j]
			if tree.height > actualHeight {
				actualHeight = tree.height
				tree.visible = true
			}
		}
	}
	return forest
}

func reverseForest(forest [][]*Tree) (res [][]*Tree) {
	res = [][]*Tree{}

	for i := len(forest) - 1; i >= 0; i-- {
		res = append(res, forest[i])
	}

	return
}

func reverseLine(line []*Tree) (res []*Tree) {
	res = []*Tree{}

	for i := len(line) - 1; i >= 0; i-- {
		res = append(res, line[i])
	}

	return
}

func printForest(forest [][]*Tree) {
	for _, line := range forest {
		for _, tree := range line {
			if tree.visible {
				fmt.Print("T")
			} else {
				fmt.Print("F")
			}

			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func totalVisibleTrees(forest [][]*Tree) int {
	total := 0
	for _, line := range forest {
		for _, tree := range line {
			if tree.visible {
				total++
			}
		}
	}
	return total
}
