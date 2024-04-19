package main

import (
	"advent/one"
	"advent/two"
	"fmt"
)

func main() {
	var err error
	var result int
	if result, err = one.PartOne(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	if result, err = one.PartTwo(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	if result, err = two.PartOne(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
