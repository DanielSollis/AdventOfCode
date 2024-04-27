package main

import (
	"advent/four"
	"advent/one"
	"advent/three"
	"advent/two"
	"fmt"
)

func main() {
	parts := []func() (int, error){
		one.PartOne,
		one.PartTwo,
		two.PartOne,
		two.PartTwo,
		three.PartOne,
		three.PartTwo,
		four.PartOne,
	}

	var err error
	var result int
	for _, part := range parts {
		if result, err = part(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
