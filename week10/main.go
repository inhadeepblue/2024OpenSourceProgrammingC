package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	//counts := 0
	var isPrime bool = true // 가독성 up, 메모리 down
	j := 2
	for j < n {
		if n%j == 0 {
			//counts = counts + 1
			isPrime = false // 더하기 연산 제거
		}
		j++
	}

	//if counts == 0 {
	if isPrime { // == 비교 연산 제거
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
