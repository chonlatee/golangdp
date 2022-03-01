package main

import "fmt"

func fib(n int, mem map[int]int) int {
	if v, ok := mem[n]; ok {
		return v
	}

	if n <= 2 {
		return 1
	}

	r := fib(n-1, mem) + fib(n-2, mem)
	mem[n] = r
	return r
}

func main() {

	mem := make(map[int]int)

	r := fib(50, mem)

	fmt.Println(r)
}
