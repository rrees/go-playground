package main

import (
	"testing"
)

func Test_diceModifiers(t *testing.T) {

	modifiers := []int{1, 2, 3, 10}
	baseDie := []int{6, 8, 12, 20}

	for _, die := range baseDie {
		for _, modifier := range modifiers {
			result := D(RollRequest{1, die, modifier})

			if (result.total < modifier || result.total > die + modifier) {
				t.Error("Invalid dice result generated: ", result.total)
			}
		}
	}
}