package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		a := make([]int, N)
		b := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&a[j], &b[j])
		}
		points := 0
		for j := 0; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if a[j] < a[k] {
					if b[j] > b[k] {
						points++
					}
				} else if b[j] < b[k] {
					points++
				}
			}
		}
		fmt.Printf("Case #%d: %d\n", i+1, points)
	}
}
