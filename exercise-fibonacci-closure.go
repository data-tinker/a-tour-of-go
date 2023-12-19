package main

import "fmt"

func fibonacci() func() int {
	num, nextNum := 0, 1
	return func() int {
		num, nextNum = nextNum, num+nextNum
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
