package main

import (
	"fmt"
	"log"
	"week11/keyboard"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j)
			j = j + 2
		}
	}
	return true
}

func main() {
	fmt.Print("Input start number : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Input end number : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
