package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var C, I int
		fmt.Scan(&C, &I)
		L := make([]int, I)
		for i := range L {
			fmt.Scan(&L[i])
		}
	Outer:
		for j := range L {
			for k := j + 1; k < len(L); k++ {
				if L[j]+L[k] == C {
					fmt.Printf("Case #%d: %d %d\n", i+1, j+1, k+1)
					break Outer
				}
			}
		}
	}
}
