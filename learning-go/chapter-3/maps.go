package main

import "fmt"

func main() {
	simpleMap()
	readWriteMap()
	commaOK()
	deleteFromMap()
	simulateSet()
}

func simpleMap() {
	fmt.Println("###Simple map###")
	score := map[string]int{"a": 2, "b": 3}
	fmt.Println(score)
}

func readWriteMap() {
	fmt.Println("###Read / write to map###")
	wins := map[string]int{}
	wins["KR"] = 1
	wins["Fylkir"] = 2

	fmt.Println("KR", wins["KR"])
	fmt.Println("Valur", wins["Valur"])
	wins["Valur"]++
	fmt.Println("Valur", wins["Valur"])
	wins["Fylkir"] = 3
	fmt.Println("Fylkir", wins["Fylkir"])
}

func commaOK() {
	fmt.Println("###Comma ok idiom###")
	m := map[string]int{
		"h": 0,
		"i": 10,
	}

	v, ok := m["h"]
	fmt.Println(v, ok)
	v, ok = m["i"]
	fmt.Println(v, ok)
	v, ok = m["!"]
	fmt.Println(v, ok)

}

func deleteFromMap() {
	fmt.Println("###Deleting from a map###")
	m := map[string]int{
		"h": 0,
		"i": 10,
	}

	var m2 map[string]int

	fmt.Println(m)
	delete(m, "i")
	fmt.Println(m)
	fmt.Println(m2)
	delete(m2, "i")
	fmt.Println(m2)
}

func simulateSet() {
	fmt.Println("###Simulating Set###")
	set := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		set[v] = true
	}

	fmt.Println("Length of vals:", len(vals), "set:", len(set))
	fmt.Println("5 in set:", set[5])
	fmt.Println("500 in set:", set[500])
	if set[100] {
		fmt.Println("100 is in the set")
	}
}
