package main

import (
	"errors"
	"fmt"
	"github.com/pbetkier/advent-of-code-2020-go/read"
	"log"
	"strconv"
	"strings"
)

func day19(input string) int {
	var rules = make([]rule, 0)

	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		log.Fatalf("cannot parse '%s', no empty line", input)
	}

	for i, s := range strings.Split(split[0], "\n") {
		split := strings.SplitN(s, ": ", 2)
		if len(split) != 2 {
			log.Fatalf("cannot parse '%s'", s)
		}

		ruleIndex, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		if ruleIndex != i {
			log.Fatalf("expected rule index '%d' to equal '%d'", ruleIndex, i)
		}

		rule, err := parseRule(split[1])
		if err != nil {
			log.Fatal(err)
		}

		rules = append(rules, rule)
	}

	fmt.Println(rules)
	return -1
}

type rule struct {
	char         byte
	subRulesAlts []subRulesAlt
}

type subRulesAlt struct {
	subRules []int
}

var emptyRule = rule{
	char:         0,
	subRulesAlts: nil,
}

func parseRule(input string) (rule, error) {
	if len(input) == 0 {
		return emptyRule, errors.New("cannot parse rule: empty")
	}

	switch input[0] {
	case '"':
		char := strings.Trim(input, "\"")
		if len(char) != 1 {
			return emptyRule, fmt.Errorf("cannot parse rule char > 1: '%s'", char)
		}
		return rule{
			char:         char[0],
			subRulesAlts: nil,
		}, nil
	default:
		var subRulesAlts = make([]subRulesAlt, 0)
		alts := strings.Split(input, " | ")
		for _, alt := range alts {
			var subRules = make([]int, 0)
			for _, subRuleStr := range strings.Split(alt, " ") {
				subRule, err := strconv.Atoi(subRuleStr)
				if err != nil {
					return emptyRule, fmt.Errorf("cannot parse '%s': %v", input, err)
				}
				subRules = append(subRules, subRule)
			}
			subRulesAlts = append(subRulesAlts, subRulesAlt{subRules: subRules})
		}
		return rule{
			char:         0,
			subRulesAlts: subRulesAlts,
		}, nil
	}
}

func main() {
	input, err := read.Text("day19/input")
	if err != nil {
		panic(err)
	}
	fmt.Println(day19(input))
}
