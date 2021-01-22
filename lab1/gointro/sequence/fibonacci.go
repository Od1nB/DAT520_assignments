package sequence

import (
	"fmt"
)

// Task 1: Fibonacci numbers
//
// fibonacci(n) returns the n-th Fibonacci number, and is defined by the
// recurrence relation F_n = F_n-1 + F_n-2, with seed values F_0=0 and F_1=1.
func fibonacci(n uint) uint {
	if n == 0 {return 0}
	var a, b, c uint = 0,1,0
	for i := 2; uint(i) <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}
