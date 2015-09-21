package main

import (
	"fmt"
)

func main() {

	var n = 100000
	var delta float64 = 1.0 / float64(n)
	var sum float64 = 0.0
	var x float64

	for i := 1; i <= n; i++ {
		x = (float64(i) - 0.5) * delta
		sum += 1.0 / (1.0 + x*x)
	}
	var pi float64 = 4.0 * delta * sum
	fmt.Println("Pi is approximately", pi, "after", n, "iterations.")
}
