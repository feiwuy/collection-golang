package main

import (
	"fmt"
)

func isPrime(N int) bool {
	ans := true
		for i := 2; i*i < N; i++ {
			if N%i == 0 {
				ans = false
			}
		}
	return ans
}

func main() {
	var ans bool
	var N int
	fmt.Scan(&N)
	ans = isPrime(N)	
	if ans {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
