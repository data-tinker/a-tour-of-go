package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
		for j := range result[i] {
			result[i][j] = uint8((i + j) / 2)
		}
	}

	return result
}

func main() {
	fmt.Println(Pic(3, 3))
}
