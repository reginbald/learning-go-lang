package main

import "fmt"

func main() {
	comparingStructs()
	convertingStructs()
}

func comparingStructs() {
	type animal struct {
		name string
		age  int
	}

	cat := animal{"stan", 10}
	cat2 := animal{"john", 20}
	cat3 := animal{"stan", 10}

	println("Is cat equal to cat2?", cat == cat2)
	println("Is cat equal to cat3?", cat == cat3)
}

func convertingStructs() {
	type person1 struct {
		name string
		age  int
	}
	type person2 struct {
		name string
		age  int
	}
	type person3 struct {
		age  int
		name string
	}
	type person4 struct {
		firstName string
		age       int
	}
	type person5 struct {
		name string
		age  int
		hair string
	}
	var a struct {
		name string
		age  int
	}

	p1 := person1{"steve", 10}
	p2 := person2(p1)
	//p3 := person3(p1) //conversion not possible
	//p4 := person4(p1) //conversion not possible
	//p5 := person5(p1) //conversion not possible
	a = p1

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("a:", a)
}
