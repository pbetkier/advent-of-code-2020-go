package main

import (
	"fmt"
	"github.com/pbetkier/advent-of-code-2020-go/read"
)

func day19(input []string) int {
	return -1
}

func main() {
	input, err := read.Lines("day19/input")
	if err != nil {
		panic(err)
	}
	fmt.Println(day19(input))
}
