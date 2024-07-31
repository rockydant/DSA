package main

import (
	"fmt"
)

func main() {
	fmt.Println(getBit(2, 5))
	fmt.Println(setBit(2, 5))
	fmt.Println(clearBit(2, 5))
}

func getBit(num int, i int) bool {
	// num:    000010
	// 1<<i:   100000
	// result: 000000
	return ((num & (1 << i)) != 0)
}

func setBit(num int, i int) int {
	// num:    000010
	// 1<<i:   100000
	// result: 100010
	return (num | (1 << i))
}

func clearBit(num int, i int) int {
	// num:    000010
	// ^(1<<i):011111
	// result: 000010
	mask := ^(1 << i)
	return num & mask
}
