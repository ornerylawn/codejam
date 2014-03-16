package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		words := strings.Split(line, " ")
		for i := 0; i < len(words)/2; i++ {
			words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
		}
		fmt.Printf("Case #%d: %s\n", i+1, strings.Join(words, " "))
	}
}
