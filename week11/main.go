package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input end number : ")
	b, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
