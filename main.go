package main

import (
	"fmt"
)

func main() {
	var err error
	var result int
	if result, err = partOne("input.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	if result, err = partTwo("input.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
