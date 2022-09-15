package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) int { return a / b }

func main() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
	}

	for _, e := range expressions {
		p1, err := strconv.Atoi(e[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := e[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator", op)
			continue
		}
		p2, err := strconv.Atoi(e[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(p1, op, p2, "=", opFunc(p1, p2))
	}
}
