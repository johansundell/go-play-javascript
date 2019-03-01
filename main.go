package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	fmt.Println(pi(1000 * 1000 * 1000))
	fmt.Println("total time taken is:", time.Since(started))
}

func pi(n int) float64 {
	f := 3.0
	for i := 1; i <= n; i++ {
		f += term(i)
	}
	return f
}

func term(n int) float64 {
	i := float64(n)*2 + 1
	if n%2 == 0 {
		return -4 / ((i - 1) * i * (i + 1))
	} else {
		return 4 / ((i - 1) * i * (i + 1))
	}
}
