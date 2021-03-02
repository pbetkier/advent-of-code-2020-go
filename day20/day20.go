package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const tileSize = 10

type tile struct {
	id     int
	top    string
	right  string
	bottom string
	left   string
}

func readTiles(filename string) ([]tile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var tiles []tile
	scanner := bufio.NewScanner(f)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
			return i + 2, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "\n")

		id, err := strconv.Atoi(split[0][len("Tile ") : len(split[0])-1])
		if err != nil {
			return nil, err
		}

		top := split[1]
		bottom := split[len(split)-1]
		var rawLeft [tileSize]uint8
		for i := 1; i < len(split); i++ {
			rawLeft[i-1] = split[i][0]
		}
		left := string(rawLeft[:])
		var rawRight [tileSize]uint8
		for i := 1; i < len(split); i++ {
			rawRight[i-1] = split[i][tileSize-1]
		}
		right := string(rawRight[:])

		tiles = append(tiles, tile{id, top, right, bottom, left})
	}

	return tiles, nil
}

func reverse(input string) string {
	output := make([]uint8, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[len(input)-1-i]
	}
	return string(output)
}

func day20(filename string) int {
	tiles, err := readTiles(filename)
	if err != nil {
		panic(err)
	}

	histogram := make(map[string]int)
	for _, t := range tiles {
		histogram[t.top] += 1
		histogram[t.right] += 1
		histogram[t.bottom] += 1
		histogram[t.left] += 1
	}

	result := 1
	for _, t := range tiles {
		topOccurrences := histogram[t.top] + histogram[reverse(t.top)]
		rightOccurrences := histogram[t.right] + histogram[reverse(t.right)]
		bottomOccurences := histogram[t.bottom] + histogram[reverse(t.bottom)]
		leftOccurrences := histogram[t.left] + histogram[reverse(t.left)]

		if isCornerTile(t, topOccurrences, rightOccurrences, bottomOccurences, leftOccurrences) {
			result *= t.id
		}
	}

	return result
}

func isCornerTile(checked tile, occurrences ...int) bool {
	ones := 0
	for _, m := range occurrences {
		switch m {
		case 1:
			ones += 1
		case 2:
			continue
		default:
			panic(fmt.Errorf("tile %v has border that doesn't occur 1 or 2 times", checked))
		}
	}

	if ones > 2 {
		panic(fmt.Errorf("tile %v has more than 2 borders that occur once", checked))
	}

	return ones == 2
}

func main() {
	result := day20("day20/input")

	fmt.Println(result)
}
