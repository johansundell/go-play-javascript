package main

import (
	"fmt"
	"time"
)

func term(n int) float64 {
	i := float64(n)*2 + 1
	if n%2 == 0 {
		return -4 / ((i - 1) * i * (i + 1))
	} else {
		return 4 / ((i - 1) * i * (i + 1))
	}
}

func GetPi(n int) float64 {
	f := 3.0
	for i := 1; i <= n; i++ {
		f += term(i)
	}
	return f
}

func main() {
	started := time.Now()
	fmt.Println(GetPi(1000 * 1000 * 1000))
	fmt.Println("total time taken is:", time.Since(started))
}
