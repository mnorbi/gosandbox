package main

import "fmt"

func sumSlice(values []int) int {
	return sum(values...)
}

func concat(s1, s2 []int) []int {
	return append(s1, s2...)
}

var (
	v  = []int{1, 2, 3, 4}
	h1 = []int{1, 2}
	h2 = []int{3, 4}
)

// Make it print out 10 10 10
// * First create a variadic function sum
// * In sumSlice call sum
// * In concat call append
func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sumSlice(v))
	fmt.Println(sumSlice(concat(h1, h2)))
}

func sum(vals ...int) int {
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}
