package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}
/*
func run(x float64) error {
	return ErrNegativeSqrt(x)
} //纱雾酱：这个run函数是多余的。考虑不创造这个run函数。
*/
func Sqrt1(x float64) (float64, error) {
	z := 1.0
	d := z
	if x <= 0 {
		return x, run(x) //纱雾酱：你应当在这里直接返回ErrNegativeSqrt类型的值，而不是返回run函数
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z) //纱雾酱：考虑函数式编程思维，你这里的写法会引起混乱
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

func Sqrt(x float64) (float64, error) {
	z := 1.0
	d := z
	if x <= 0 {
		return x, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		z =z- (z*z - x) / (2 * z)
		if math.Abs(d-z) < 0.00000001 {
			break
		}

	}
	return z, nil
}
