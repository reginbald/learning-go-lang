package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modifyMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	i := 2
	s := "Hello"
	p := person{}
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	fmt.Println("---------")

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modifyMap(m)
	fmt.Println(m)

	sl := []int{1, 2, 3}
	modifySlice(sl)
	fmt.Println(sl)
}
