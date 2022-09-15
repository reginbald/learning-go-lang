package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(namedReturnValues(5, 2))
}

func namedReturnValues(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err

	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, nil
}
