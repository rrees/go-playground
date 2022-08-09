package main

import (
	"fmt"
)

func main() {
	fmt.Print("|")
	for i := 0; i < 20; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(D(RollRequest{2, 6, 0}))
	}
	fmt.Println("|")
}
