package main

import "fmt"

var memoCache = make(map[int]int)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	if v, ok := memoCache[n]; ok {
		return v
	}

	memoCache[n] = fibonacci(n-1) + fibonacci(n-2)

	return memoCache[n]
}

func main() {
	fmt.Println(fibonacci(40))
}
