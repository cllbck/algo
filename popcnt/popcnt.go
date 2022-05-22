package main

import (
	"fmt"
	"math/bits"
)

func popCount(number uint) int {
	count := 0
	for number > 0 {
		if number&1 == 1 {
			count++
		}
		number >>= 1
	}
	return count
}

func popCount2(number uint) int {
	count := 0
	for number > 0 {
		count++
		number &= number - 1
	}
	return count
}

var popCountMap = map[uint]int{}

func initPopCountMap() {
	for i := 0; i <= 255; i++ {
		popCountMap[uint(i)] = bits.OnesCount(uint(i))
	}
}

func popCount3(number uint) int {
	count := 0
	for number > 0 {
		count += popCountMap[number&255]
		number >>= 8
	}
	return count
}

func main() {
	var n uint = 345634576362345
	fmt.Printf("counts 1 bit for %v - %v \n", n, bits.OnesCount(n))
	fmt.Printf("counts 1 bit for %v - %v \n", n, popCount(n))
	fmt.Printf("counts 1 bit for %v - %v \n", n, popCount2(n))
	fmt.Printf("counts 1 bit for %v - %v \n", n, popCount3(n))
}

func init() {
	initPopCountMap()
}
