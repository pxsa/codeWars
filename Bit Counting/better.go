package main 

import (
	"math/bits"
)

func Solution2 (num uint) int {
	return bits.OnesCount(num)
}

func Solution3 (num uint) int {
	var count int
	for num>0 {
		if (num & 1 == 1) {
			count++
		}
		num = num >> 1
	}
	return count 
}