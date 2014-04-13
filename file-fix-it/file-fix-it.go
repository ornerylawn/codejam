package main

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N, M int
		fmt.Scan(&N, &M)
		dirs := map[string]bool{}
		for j := 0; j < N; j++ {
			var line string
			fmt.Scan(&line)
			dirs[line] = true
		}
		count := 0
		for j := 0; j < M; j++ {
			var line string
			fmt.Scan(&line)
			segments := strings.Split(line, "/")[1:]
			for i := range segments {
				dir := "/" + strings.Join(segments[:i+1], "/")
				if dirs[dir] {
					continue
				}
				dirs[dir] = true
				count++
			}
		}
		fmt.Printf("Case #%d: %d\n", i+1, count)
	}
}
