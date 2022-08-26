package loopsAndFunctions

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, zn, d, delta := x/2, 0.0, 0.0, 0.000001
	for {
		zn -= (z*z - x) / (2 * z)
		fmt.Println(zn)
		d = z - zn
		if d < 0.0 && d > delta || d > 0.0 && d < delta {
			return zn
		}
		z = zn
	}
}

func Run() {
	fmt.Println("Sqrt(2): ", Sqrt(2))
	fmt.Println("math.Sqrt(2): ", math.Sqrt(2))
}
