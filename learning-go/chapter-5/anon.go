package main

import (
	"fmt"
	"sort"
)

func anonFuncs() {
	fmt.Println(func(a, b int) int { return a + b }(1, 2))
	outerVar := 0
	for i := 0; i < 5; i++ {
		func(j int) {
			outerVar++
			fmt.Println("Index: ", j, "Outer variable:", outerVar)
		}(i)
	}
}

func passingFunctions() {
	numbers := []int{4, 3, 2, 1}
	fmt.Println("Before:", numbers)
	sort.Slice(numbers, func(a, b int) bool { return a > b })
	fmt.Println("After:", numbers)
}

func main() {
	anonFuncs()
	fmt.Println("------------")
	passingFunctions()
}
