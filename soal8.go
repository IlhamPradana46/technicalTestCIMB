package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	var res []int
	for i := 0; i < len(expression); i++ {
		oper := expression[i]
		if oper == '+' || oper == '-' || oper == '*' {
			s1 := diffWaysToCompute(expression[:i])
			s2 := diffWaysToCompute(expression[i+1:])

			for _, elS1 := range s1 {
				for _, elS2 := range s2 {
					if oper == '+' {
						res = append(res, elS1+elS2)
					} else if oper == '-' {
						res = append(res, elS1-elS2)
					} else if oper == '*' {
						res = append(res, elS1*elS2)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		expInt, _ := strconv.Atoi(expression)
		res = append(res, expInt)
	}
	return res
}

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
