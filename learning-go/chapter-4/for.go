package main

import "fmt"

func main() {
	noContinue()
	withContinue()
	forRangeSlices()
	forRangeIgnoreIndex()
	forRangeIgnoreValue()
	forRangeMaps()
	forRangeStrings()
	forRangeValueCopy()
	labels()
}

func noContinue() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func withContinue() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

func forRangeSlices() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func forRangeIgnoreIndex() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func forRangeIgnoreValue() {
	uniqueNames := map[string]bool{"Fred": true, "Sam": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

func forRangeMaps() {
	m := map[string]int{"a": 1, "c": 3, "b": 2}
	for i := 0; i < 3; i++ {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func forRangeStrings() {
	samples := []string{"hello", "apple_π"}
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func forRangeValueCopy() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

func labels() {
	samples := []string{"hello", "world"}
outer:
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				fmt.Println()
				continue outer
			}
		}
	}
}
