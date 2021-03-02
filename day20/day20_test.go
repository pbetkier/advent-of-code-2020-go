package main

import "testing"

func TestReadTiles(t *testing.T) {
	tiles, err := readTiles("input-example")

	if err != nil {
		t.Error(err)
	}

	if len(tiles) != 9 {
		t.Errorf("len(tiles) want %d, got %d", 9, len(tiles))
	}

	want := tile{1951, "#.##...##.", ".#####..#.", "#...##.#..", "##.#..#..#"}
	got := tiles[1]
	if got != want {
		t.Errorf("tile[1] want %v, got %v", want, got)
	}
}

func TestPart1(t *testing.T) {
	got := day20("input-example")
	want := 20899048083289

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
