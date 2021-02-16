package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay19Part1(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{`0: 1 2
1: "a"
2: 1 3
3: "b"

aab
aba
abb`, 1},
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
			if got := day19(c.input); got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}

func TestCartesianProduct(t *testing.T) {
	cases := []struct {
		sets [][]string
		want []string
	}{
		{
			[][]string{{"a", "b"}, {"1", "2"}},
			[]string{"a1", "a2", "b1", "b2"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("solves #%d", i), func(t *testing.T) {
			if diff := cmp.Diff(c.want, cartesianProduct(c.sets...)); diff != "" {
				t.Errorf("cartesianProduct(...) mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
