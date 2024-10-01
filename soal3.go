package main

func getSmallestString(n int, k int) string {
	result := make([]byte, n)
	angka := k - n

	for i := n - 1; i >= 0; i-- {
		if angka > 25 {
			angka -= 25
			result[i] = 'z'
		} else {
			result[i] = byte(angka + 'a')
			angka -= angka
		}
	}
	return string(result)
}
