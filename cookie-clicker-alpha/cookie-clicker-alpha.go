package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var C, F, X float64
		fmt.Scan(&C, &F, &X)
		rate := 2.0
		t := 0.0
		for {
			t1 := X / rate
			tfarm := C / rate
			t2 := tfarm + (X / (rate + F))
			if t1 <= t2 {
				fmt.Printf("Case #%d: %.7f\n", i+1, t+t1)
				break
			}
			t += tfarm
			rate += F
		}
	}
}
