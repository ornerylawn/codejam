package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var L, D, N int
	fmt.Scan(&L, &D, &N)
	words := make([]string, D)
	for i := 0; i < D; i++ {
		fmt.Scan(&words[i])
	}
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		s.Scan()
		pat := s.Text()
		count := 0
		for _, word := range words {
			if match(word, pat) {
				count++
			}
		}
		fmt.Printf("Case #%d: %d\n", i+1, count)
	}
}

func match(s, pat string) bool {
	for i := 0; i < len(s); i++ {
		if string(pat[0]) != "(" {
			if s[i] != pat[0] {
				return false
			}
			pat = pat[1:]
			continue
		}
		found := false
		for pat = pat[1:]; string(pat[0]) != ")"; pat = pat[1:] {
			if s[i] == pat[0] {
				found = true
			}
		}
		pat = pat[1:]
		if !found {
			return false
		}
	}
	return true
}
