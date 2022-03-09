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

func fibtabu(n int) int {
	f := make([]int, n+3, n+3)
	f[1] = 1
	for i := 0; i <= n; i++ {
		f[i+1] += f[i]
		f[i+2] += f[i]
	}

	return f[n]
}

func main() {

	mem := make(map[int]int)

	r := fib(50, mem)

	fmt.Println(r)

	fmt.Println(fibtabu(50))
}
