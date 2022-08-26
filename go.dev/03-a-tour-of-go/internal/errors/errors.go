package errors

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z, zn, d, delta := x/2, 0.0, 0.0, 0.000001
	for {
		zn -= (z*z - x) / (2 * z)
		d = z - zn
		if d < 0.0 && d > delta || d > 0.0 && d < delta {
			return zn, nil
		}
		z = zn
	}
}

func Run() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
