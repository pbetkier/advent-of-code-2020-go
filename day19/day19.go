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
	rules, messages := parse(input)

	valid := toSet(validMessages(rules, 0))

	validCount := 0
	for _, m := range messages {
		if _, ok := valid[m]; ok {
			validCount += 1
		}
	}
	return validCount
}

func toSet(keys []string) map[string]struct{} {
	var result = make(map[string]struct{})
	for _, c := range keys {
		result[c] = struct{}{}
	}
	return result
}

func validMessages(rules []rule, nextRuleIndex int) []string {
	nextRule := rules[nextRuleIndex]

	if nextRule.char != 0 {
		return []string{string(nextRule.char)}
	}

	var result = make([]string, 0)

	for _, alt := range nextRule.subRulesAlts {
		var messagePartsOptions = make([][]string, 0)
		for _, sr := range alt.subRules {
			messagePartsOptions = append(messagePartsOptions, validMessages(rules, sr))
		}

		result = append(result, cartesianProduct(messagePartsOptions...)...)
	}

	return result
}

func cartesianProduct(sets ...[]string) []string {
	return doCartesianProduct(sets, "")
}

func doCartesianProduct(sets [][]string, prefix string) []string {
	if len(sets) == 0 {
		return []string{prefix}
	}

	var result = make([]string, 0)
	for _, s := range sets[0] {
		result = append(result, doCartesianProduct(sets[1:], prefix+s)...)
	}

	return result
}

func parse(input string) ([]rule, []string) {
	var messages = make([]string, 0)

	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		log.Fatalf("cannot parse '%s', no empty line", input)
	}

	ruleLines := strings.Split(split[0], "\n")
	var rules = make([]rule, len(ruleLines))
	for _, s := range ruleLines {
		split := strings.SplitN(s, ": ", 2)
		if len(split) != 2 {
			log.Fatalf("cannot parse '%s'", s)
		}

		ruleIndex, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}

		rule, err := parseRule(split[1])
		if err != nil {
			log.Fatal(err)
		}

		rules[ruleIndex] = rule
	}

	for _, s := range strings.Split(split[1], "\n") {
		messages = append(messages, s)
	}

	return rules, messages
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
