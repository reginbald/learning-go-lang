package main

import "fmt"

func main() {
	sharedMemory()
	overwriteScenario()
	limitedCapacity()
}

func sharedMemory() {
	fmt.Println("Subslices share memory")
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]

	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))

	y[1] = 20

	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))
}

func overwriteScenario() {
	fmt.Println("Subslice overwrite scenario:")
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)

	y := x[:2]
	z := x[2:]
	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))
	fmt.Println("z:", z, cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))
	fmt.Println("z:", z, cap(z))
}

func limitedCapacity() {
	fmt.Println("Subslice with limited parent capacity:")
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]

	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))

	x = append(x, 10)
	y = append(y, 20)

	fmt.Println("x:", x, cap(x))
	fmt.Println("y:", y, cap(y))
}
