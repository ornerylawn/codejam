package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		naomi := make([]float64, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&naomi[j])
		}
		ken := make([]float64, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&ken[j])
		}
		sort.Float64s(naomi)
		sort.Float64s(ken)
		usedKen := make([]bool, N)
		war := 0
		for j := 0; j < N; j++ {
			x := naomi[j]
			win := true
			for k := 0; k < N; k++ {
				if !usedKen[k] && ken[k] > x {
					win = false
					usedKen[k] = true
					break
				}
			}
			if win {
				war++
			}
		}
		deceit := 0
		usedNaomi := make([]bool, N)
		for j := 0; j < N; j++ {
			x := ken[j]
			win := false
			for k := 0; k < N; k++ {
				if !usedNaomi[k] && naomi[k] > x {
					win = true
					usedNaomi[k] = true
					break
				}
			}
			if win {
				deceit++
			}
		}
		fmt.Printf("Case #%d: %d %d\n", i+1, deceit, war)
	}
}
