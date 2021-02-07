package main

import (
	"fmt"
	"testing"
)

func TestDay18Part1(t *testing.T) {
	cases := []struct {
		expr string
		want int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("solves '%s'", c.expr), func(t *testing.T) {
			if got := day18Part1SingleExpr(c.expr); got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}

func TestDay18Part2(t *testing.T) {
	cases := []struct {
		expr string
		want int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
		{"1 * 2 * (3 + 4) + 5 * 6", 144},
		{"(2 * (3 * 4) + 1) + 1", 27},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("solves '%s'", c.expr), func(t *testing.T) {
			if got := day18Part2SingleExpr(c.expr); got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}
