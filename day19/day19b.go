package main

import (
	"fmt"
	"github.com/pbetkier/advent-of-code-2020-go/read"
)

func day19b(input string) int {
	rules, messages := parse(input)
	return validMessagesB(messages, rules)
}

func validMessagesB(messages []string, rules map[int]rule) int {
	// 0: 8 11
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// results in:
	// 42 42 42* 42x 31x 31

	options42 := validMessages(rules, 42)
	options31 := validMessages(rules, 31)
	if len(options42[0]) != len(options31[0]) {
		panic(fmt.Errorf("expected len(42) == len(31), was %d != %d", len(options42[0]), len(options31[0])))
	}

	options42Set := toSet(options42)
	options31Set := toSet(options31)

	validCount := 0
	for _, m := range messages {
		if isValidB(m, options42Set, options31Set, len(options42[0])) {
			validCount += 1
		}
	}

	return validCount
}

func isValidB(message string, options42 map[string]struct{}, options31 map[string]struct{}, optionLen int) bool {
	segments := len(message) / optionLen
	if segments*optionLen != len(message) || segments < 3 {
		return false
	}

	for n4231 := 0; n4231*2 <= segments-3; n4231 += 1 {
		n31 := n4231 + 1
		n42 := segments - n31

		ok, messageRest := startsWithTimesN(message, n42, options42, optionLen)
		if !ok {
			continue
		}
		ok, messageRest = startsWithTimesN(messageRest, n31, options31, optionLen)
		if !ok {
			continue
		}
		if len(messageRest) == 0 {
			return true
		}
	}

	return false
}

func startsWith(message string, options map[string]struct{}, optionsLen int) bool {
	if len(message) < optionsLen {
		return false
	}

	if _, ok := options[message[:optionsLen]]; ok {
		return true
	}
	return false
}

func startsWithTimesN(message string, n int, options map[string]struct{}, optionsLen int) (bool, string) {
	for i := 0; i < n; i += 1 {
		if !startsWith(message, options, optionsLen) {
			return false, message
		}
		message = message[optionsLen:]
	}

	return true, message
}

func mainB() {
	input, err := read.Text("day19/input-part2")
	if err != nil {
		panic(err)
	}

	fmt.Println(day19b(input))
}
