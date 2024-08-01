package main

import "fmt"

type memoized struct {
	f     func(int) int
	cache map[int]int
}

func memoize(f func(int) int) *memoized {
	return &memoized{f, make(map[int]int)}
}

func (m *memoized) call(x int) int {
	if v, ok := m.cache[x]; ok {
		return v
	}

	result := m.f(x)

	m.cache[x] = result
	return result
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	m := memoize(fibonacci)
	fmt.Println(m.call(5))
}
