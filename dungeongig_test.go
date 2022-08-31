package main

import (
	"testing"
)

func Test_allDifferentScores(t *testing.T) {
	result := CalculateStat([3]int{1, 2, 3})

	if result != 2 {
		t.Error("Stat was expected to be 2 but got ", result)
	}
}


func Test_allSameScores(t *testing.T) {
	result := CalculateStat([3]int{4,4,4})

	if result != 4 {
		t.Error("Expected the Stat to be 4 but got ", result)
	}
}

func Test_doubledScores(t *testing.T) {
	type doubleScores struct {
		scores [3]int
		expectedStat int
	}

	testDataset := []doubleScores{
		{scores: [3]int{2, 2, 5}, expectedStat: 5},
		{scores: [3]int{1, 4, 4}, expectedStat: 4},
	}

	for _, testData := range testDataset {
		result := CalculateStat(testData.scores)

		if result != testData.expectedStat {
			t.Error("Expected to get ", testData.expectedStat, " but got ", result)
		}
	}
}