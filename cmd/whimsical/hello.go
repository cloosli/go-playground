package main

import (
	"fmt"
	"math/rand"
)

func main() {
	t := "Hello World!"
	s := []rune(t)

	for {
		rand.Shuffle(len(s), func(i int, j int) {
			s[i], s[j] = s[j], s[i]
		})
		fmt.Println(string(s))
		if string(s) == t {
			break
		}
	}
}
