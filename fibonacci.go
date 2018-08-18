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
	return func(x int) int { //紗霧：这个练习的意思是让fibonacci生成一个f()函数，这个f()每执行一次就返回下一个值，也就是说这个函数每次执行返回的值都不一样。这个函数并不需要你传入一个x。当然实际上，你的想法更加正确，只是这个练习的目的是让你理解闭包。在这个练习之外，我们一般不推荐让函数每次执行都返回不同的值。因此你return的应该是func() int，而不是你现在的func(x int) int
		if x < 2 {
			return x
		}
		mid, ans = ans, ans+mid
		return ans
	}
}

func fibonacci2() func() int {

	mid, mid1, ans := 0, 1, 0
	return func() int {
		ans = mid
		mid, mid1 = mid1, mid+mid1
		return ans
	}
}
