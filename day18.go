package main

import (
	"bufio"
	"fmt"
	"os"
)

func day18Part1() int {
	return day18(day18Part1SingleExpr)
}

func day18Part2() int {
	return day18(day18Part2SingleExpr)
}

func day18(computeSingle func(string) int) int {
	exprs, err := readDay18("testdata/day18-input")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, e := range exprs {
		result += computeSingle(e)
	}
	return result
}

func readDay18(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var text []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}

func day18Part1SingleExpr(expr string) int {
	return doDay18Part1SingleExpr(expr, new(int))
}

func doDay18Part1SingleExpr(expr string, i *int) int {
	acc := 0
	op := add

	for ; *i < len(expr); *i += 1 {
		switch char := expr[*i]; char {
		case ' ':
			continue
		case '+':
			op = add
		case '*':
			op = mul
		case '(':
			*i += 1
			acc = op(acc, doDay18Part1SingleExpr(expr, i))
		case ')':
			return acc
		default:
			acc = op(acc, int(char-'0'))
		}
	}

	return acc
}

func day18Part2SingleExpr(expr string) int {
	return doDay18Part2SingleExpr(expr, new(int))
}

func doDay18Part2SingleExpr(expr string, i *int) int {
	acc := 0
	op := uint8('+')

	for ; *i < len(expr); *i += 1 {
		switch char := expr[*i]; char {
		case ' ':
			continue
		case '+', '*':
			op = char
		case '(':
			*i += 1
			switch op {
			case '+':
				acc = add(acc, doDay18Part2SingleExpr(expr, i))
			case '*':
				acc = mul(acc, doDay18Part2SingleExpr(expr, i))
			}
		case ')':
			return acc
		default:
			switch op {
			case '+':
				acc = add(acc, int(char-'0'))
			case '*':
				acc = mul(acc, doDay18Part2SingleExpr(expr, i))
			}
		}
	}

	return acc
}

func main() {
	fmt.Println(day18Part1())
	fmt.Println(day18Part2())
}
