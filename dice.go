package main

import (
	"math/rand"
	"time"
)

func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

type RollResult struct {
	total int
	results []int
}

func D(numberOfDice int, maxSides int) RollResult {
	r := random()
	results := make([]int, numberOfDice)
	for i := 0; i < numberOfDice; i++ {
		results[i] = 1 + r.Intn(maxSides)
	}

	var total = 0

	for i := 0; i < len(results); i++ {
		total += results[i]
	}

	return RollResult{
		total,
		results,
	}
}
