package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]

	numberOfDice := 2

	if(len(arguments) > 0) {
		convertedNumber, err := strconv.Atoi(arguments[0])
		if err == nil {
			numberOfDice = convertedNumber
		}
	}

	rolls := make([]RollResult, numberOfDice)

	for i := 0; i < numberOfDice; i++ {
		rolls[i] = D(RollRequest{2, 3, -3})
	}

	fmt.Println(rolls)
}
