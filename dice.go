package main

import (
	"math/rand"
	"time"
)

func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func D(numberOfDice int, maxSides int) int {
	r := random()
	results := make([]int, numberOfDice)
	for i := 0; i < numberOfDice; i++ {
		results[i] = 1 + r.Intn(maxSides)
	}

	var total = 0

	for i := 0; i < len(results); i++ {
		total += results[i]
	}

	return total
}
