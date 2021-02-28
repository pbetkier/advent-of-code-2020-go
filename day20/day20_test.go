package day20

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
