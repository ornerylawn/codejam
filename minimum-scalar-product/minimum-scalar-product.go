package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		b := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&b[j])
		}
		sort.Ints(a)
		sort.Ints(b)
		sum := 0
		for j := 0; j < n; j++ {
			sum += a[j] * b[n-1-j]
		}
		fmt.Printf("Case #%d: %d\n", i+1, sum)
	}
}
