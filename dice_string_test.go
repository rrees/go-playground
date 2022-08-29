package main

import (
	"testing"
)

func Test_diceStringNormalisationForMissingModifiers(t *testing.T) {
	inputString := "1d6"
	result := normaliseString(inputString)

	if result != "1d6+0" {
		t.Error("Dice string ", inputString, " not normalised, got ", result)
	}
}

func Test_diceStringNormalisationForProvidedModifiers(t *testing.T) {
	inputString := "1d6+2"
	result := normaliseString(inputString)

	if result != "1d6+2" {
		t.Error("Dice string ", inputString, " not normalised, got ", result)
	}
}
