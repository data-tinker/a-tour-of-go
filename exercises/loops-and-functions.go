package main

import (
	"fmt"
	"math"
)

func isConverged(a float64, b float64, d float64) bool {
	return math.Abs(a-b) < d
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	delta := 0.000001
	for {
		tmp = z - (z*z-x)/(2*z)
		if isConverged(z, tmp, delta) {
			break
		}
		z = tmp
	}
	return tmp
}

func main() {
	fmt.Println(Sqrt(100))
	fmt.Println(math.Sqrt(100))
}
