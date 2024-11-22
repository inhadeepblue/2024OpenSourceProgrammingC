package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args[1:], len(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices[1])
	for _, slice := range slices {
		fmt.Println(slice)
	}
	slices = append(slices, "forever", "!")
	fmt.Println(slices, len(slices))
}
