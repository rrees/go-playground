package main

import (
	"regexp"
)

var NoModiferRegExp, _ = regexp.Compile(`^(\d)+d(\d)+$`)

func normaliseString(diceString string) string {
	if(NoModiferRegExp.MatchString(diceString)) {
		return diceString + "+0"
	}

	return diceString
}