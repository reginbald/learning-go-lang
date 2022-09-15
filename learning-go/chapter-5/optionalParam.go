package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	age       int
}

func optionalNamedParamsSimulation(opts MyFuncOpts) error {
	fmt.Println("Name:", opts.FirstName, opts.LastName)
	return nil
}

func main() {
	optionalNamedParamsSimulation(MyFuncOpts{LastName: "Smith"})
}
