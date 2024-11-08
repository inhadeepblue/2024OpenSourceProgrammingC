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

func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	fmt.Print("Input start number : ")
	n1 := getInteger()
	fmt.Print("Input end number : ")
	n2 := getInteger()

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
