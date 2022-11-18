package main 

import (
	"fmt"
)

func CountBits(num uint) int {
	var count int
	for num != 0 {
		if res := num%2; res == 1 {
			count++
		}
		num = num/2
	}
	return count
}



func main() {
	fmt.Println(CountBits(10))
}