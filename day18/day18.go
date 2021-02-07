package main

import (
	"fmt"
	"github.com/pbetkier/advent-of-code-2020-go/read"
)

func day18Part1() int {
	return day18(day18Part1SingleExpr)
}

func day18Part2() int {
	return day18(day18Part2SingleExpr)
}

func day18(computeSingle func(string) int) int {
	exprs, err := read.Lines("day18/input")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, e := range exprs {
		result += computeSingle(e)
	}
	return result
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
		switch c := expr[*i]; c {
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
			acc = op(acc, int(c-'0'))
		}
	}

	return acc
}

func day18Part2SingleExpr(expr string) int {
	return doDay18Part2SingleExpr(expr, new(int))
}

func doDay18Part2SingleExpr(expr string, i *int) int {
	acc := 0

	for *i < len(expr) {
		c := expr[*i]
		*i += 1
		switch c {
		case ' ', '+':
			continue
		case '*':
			acc *= doDay18Part2SingleExpr(expr, i)
			if expr[*i-1] == ')' {
				*i -= 1 // let '(' subroutine consume its matching ')'
			}
		case '(':
			acc += doDay18Part2SingleExpr(expr, i)
		case ')':
			return acc
		default:
			acc += int(c - '0')
		}
	}

	return acc
}

func main() {
	fmt.Println(day18Part1())
	fmt.Println(day18Part2())
}
