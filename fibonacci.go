package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		if x < 2 {
			return x
		}
		a := fibonacci()
		//不能用f
		return (a(x-1) + a(x-2))
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(i), " ")
	}
}

func fibonacci1() func(int) int {
	mid, ans := 0, 1
	return func(x int) int {
		if x < 2 {
			return x
		}
		mid, ans = ans, ans+mid
		return ans
	}
}
