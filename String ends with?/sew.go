package main

import (
	"fmt"
)

func solution(str, ending string) bool {
	// Your code here!
	lenEnd := len(ending)
	lenStr := len(str)

	if lenStr < lenEnd {
		return false
	}

	str_new := str[lenStr-lenEnd:]
	return str_new == ending
}

func main() {
	fmt.Println(solution("", ""))
	fmt.Println(solution(" ", ""))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("banana", "ana"))
	fmt.Println(solution("a", "z"))
	fmt.Println(solution("", "t"))
	fmt.Println(solution("$a = $b + 1", "+1"))
	fmt.Println(solution("    ", "   "))
	fmt.Println(solution("1oo", "100"))
	
}


