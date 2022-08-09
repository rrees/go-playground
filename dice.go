package main

import (
	"math/rand"
	"time"
)

func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

type RollRequest struct {
	numberOfDice int
	sides        int
	modifier     int
}

type RollResult struct {
	total   int
	results []int
}

func D(roll RollRequest) RollResult {
	r := random()
	results := make([]int, roll.numberOfDice)
	
	for i := 0; i < roll.numberOfDice; i++ {
		results[i] = 1 + r.Intn(roll.sides)
	}

	var total = 0

	for i := 0; i < len(results); i++ {
		total += results[i]
	}

	return RollResult{
		total + roll.modifier,
		results,
	}
}
