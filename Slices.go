package main

import (
	"fmt"
)


func fibonacci() func() int {
	var x, prev1, prev2 int

	return func() int {
		var fib int
		var tmp int
		x += 1

		if x == 1 {
			fib = 0
			prev1 = fib
		} else if x == 2 {
			fib = 1
			tmp = fib
			prev1 = tmp
		} else {
			fib = prev1 + prev2
			tmp = prev1
			prev1 = fib
			prev2 = tmp
		}

		return fib
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 25; i++ {
		fmt.Println(f())
	}



}

