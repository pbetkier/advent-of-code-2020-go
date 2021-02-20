package main

import (
	"fmt"
	"log"
	"net/http"
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
			if got := day19a(c.input); got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}

func TestDay19Part2(t *testing.T) {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	input := `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31 | 42 11 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42 | 42 8
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`
	want := 12

	if got := day19b(input); got != want {
		t.Errorf("got %d; want %d", got, want)
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
