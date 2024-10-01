package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	say := countAndSay(n - 1)
	var result strings.Builder

	for i := 0; i < len(say); i++ {
		ch := string(say[i])
		count := 1

		for i < len(say)-1 && say[i] == say[i+1] {
			count++
			i++
		}

		result.WriteString(strconv.Itoa(count) + ch)

	}

	return result.String()
}
