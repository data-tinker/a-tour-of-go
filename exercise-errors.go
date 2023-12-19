package main

import (
	"fmt"
	"log"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	result, _ := Sqrt(2)
	log.Println(result)

	result, err := Sqrt(-2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
