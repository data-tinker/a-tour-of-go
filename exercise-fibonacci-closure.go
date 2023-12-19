package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevNum, currentNum, nextNum := 0, 0, 1
	return func() int {
		prevNum, currentNum, nextNum = currentNum, nextNum, currentNum+nextNum
		return prevNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
