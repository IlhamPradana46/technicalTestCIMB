package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	resultMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			el := board[i][j]

			if el != '.' {
				rowKey := string(el) + "row" + strconv.Itoa(i)
				colKey := string(el) + "col" + strconv.Itoa(j)
				boxKey := string(el) + "subbox" + strconv.Itoa(i/3) + strconv.Itoa(j/3)

				if resultMap[rowKey] || resultMap[colKey] || resultMap[boxKey] {
					return false
				}

				resultMap[rowKey] = true
				resultMap[colKey] = true
				resultMap[boxKey] = true
			}
		}
	}

	return true
}

func main() {
	arr := [][]byte{
		{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
		{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(arr))
}
