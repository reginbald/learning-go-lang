package main

import "fmt"

func main() {
	sliceExpression()
	intToString()
	stringToBytesOrRunes()
}

func sliceExpression() {
	fmt.Println("Extracting substring")
	var s string = "Hello"
	var s2 = s[:2]

	fmt.Println(s2)
}

func intToString() {
	fmt.Println("Converting int to string")
	var x int = 65
	var s = string(x)
	fmt.Println(s)
}

func stringToBytesOrRunes() {
	var s string = "Hello ✌"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)

	fmt.Println("String:", s)
	fmt.Println("Bytes:", bs)
	fmt.Println("Runes:", rs)
}
