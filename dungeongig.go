package main

import (
	"github.com/rrees/gobatteries/math"
)

func CalculateStat(scores [3]int) int {
	if scores[0] != scores[1] && scores[1] != scores[2] {
		return scores[1]
	}
	return math.MaxInt(scores[1], scores[2])
}
