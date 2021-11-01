package main

import "fmt"

func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func printFib(n int, done chan struct{}) {
	fmt.Println(n, ":", fib(n))
	done <- struct{}{}
}

func main() {
	d := make(chan struct{})
	n := 10
	for i := 0; i < n; i++ {
		go printFib(i, d)
	}
	for i := 0; i < n; i++ {
		<-d
	}
}
