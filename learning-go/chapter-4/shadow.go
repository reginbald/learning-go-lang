package main

import "fmt"

func main() {
	simpleShadowing()
	multiAssignShadowing()
}

func simpleShadowing() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func multiAssignShadowing() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 10
		fmt.Println(x, y)
	}
	fmt.Println(x)
}
