package main

import (
	"fmt"
	"strconv"
)

func removeLastIndex(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

func calPoints(operations []string) int {
	var arrayResult []int
	var result int

	for _, e := range operations {
		if e == "C" {
			if len(arrayResult) > 0 {
				arrayResult = arrayResult[:len(arrayResult)-1]
			}
		} else if e == "D" {
			if len(arrayResult) > 0 {
				arrayResult = append(arrayResult, arrayResult[len(arrayResult)-1]*2)
			}
		} else if e == "+" {
			if len(arrayResult) > 1 {
				arrayResult = append(arrayResult, arrayResult[len(arrayResult)-2]+arrayResult[len(arrayResult)-1])
			}
		} else {
			insertInt, err := strconv.Atoi(e)
			if err == nil {
				arrayResult = append(arrayResult, insertInt)
			}
		}
	}

	for _, e := range arrayResult {
		result += e
	}
	fmt.Println(result)

	return result
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	calPoints(ops)
}
