package main

import (
	"fmt"
	"math"
)

func longestValidParentheses(s string) int {
	var count int
	arrayChar := []int{-1}

	for i, c := range s {
		if string(c) == "(" {
			arrayChar = append(arrayChar, i)
		} else {
			arrayChar = arrayChar[:len(arrayChar)-1]
			if len(arrayChar) == 0 {
				arrayChar = append(arrayChar, i)
			} else {
				count = int(math.Max(float64(count), float64(i-arrayChar[len(arrayChar)-1])))
			}
		}
	}
	return count

}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
