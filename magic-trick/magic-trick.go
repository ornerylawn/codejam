package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var a int
		fmt.Scan(&a)
		gridA := make([][]int, 4)
		for j := 0; j < 4; j++ {
			gridA[j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				fmt.Scan(&gridA[j][k])
			}
		}
		var b int
		fmt.Scan(&b)
		gridB := make([][]int, 4)
		for j := 0; j < 4; j++ {
			gridB[j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				fmt.Scan(&gridB[j][k])
			}
		}
		count := 0
		card := 0
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if gridA[a-1][j] == gridB[b-1][k] {
					count++
					card = gridA[a-1][j]
				}
			}
		}
		switch count {
		case 0:
			fmt.Printf("Case #%d: Volunteer cheated!\n", i+1)
		case 1:
			fmt.Printf("Case #%d: %d\n", i+1, card)
		default:
			fmt.Printf("Case #%d: Bad magician!\n", i+1)
		}
	}
}
