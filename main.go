package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomSeq := randSeq(45)

	tick := time.Tick(5 * time.Second)
	for range tick {
		fmt.Printf("%v %v\n", time.Now().UTC(), randomSeq)
	}
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}