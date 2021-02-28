package day20

import (
	"bufio"
	"bytes"
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
