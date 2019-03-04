package main

import (
	"fmt"
	"time"

	"github.com/johansundell/go-play-javascript/pkg/mypi"
)

func main() {
	started := time.Now()
	fmt.Println(mypi.GetPi(1000 * 1000 * 1000))
	fmt.Println("total time taken is:", time.Since(started))
}
