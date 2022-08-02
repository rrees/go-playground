package main

import (
	"math/rand"
	"time"
)

func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func D(max_sides int) int {
	return 1 + random().Intn(max_sides)
}
