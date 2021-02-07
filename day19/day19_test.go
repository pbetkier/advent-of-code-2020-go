package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay19Part1(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"

aab
aba
abb`, 2},
		{`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`, 2},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("solves #%d", i), func(t *testing.T) {
			if got := day19(strings.Split(c.input, "\n")); got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}
