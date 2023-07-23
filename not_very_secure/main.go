package main

import (
	"fmt"
)

func main () {
	str := "45464dfjsgksfJFSGLSDJFKL"
	fmt.Println(validation(str))
}

func validation(str string) bool {
	if len(str) <= 0{
		return false
	}
	  
	for _, char := range str {
		if !((48 <= char && char <= 57) || (65 <= char && char <= 90) || (97 <= char && char <= 122)) {
			return false
		}
	}
	
	return true
}

// better solution

// pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
// return pattern.MatchString(str)