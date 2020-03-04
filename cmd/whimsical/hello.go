package main

import (
	"fmt"
	"math/rand"
)

func main() {
	t := "Hello World!"
	c := make(chan string)

	for i := 0; i < 6; i++ {
		go gopher(t, c)
	}

	for s := range c {
		fmt.Println(s)
		if s == t {
			break
		}
	}
}

func gopher(t string, c chan string) {
	s := []rune(t)
	for {
		rand.Shuffle(len(s), func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		c <- string(s)
	}
}
