package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func run(x float64) error {
	return ErrNegativeSqrt(x)
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	d := z
	if x <= 0 {
		return x, run(x)
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(d-z) < 0.00000001 {
			break
		}

	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
