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
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                // 줄바꿈등 제거. 파이썬의 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 10, 32) // 10진수 정수형(32bit)으로 변환

	var grade string
	if score >= 90 {
		grade = "A"
	} else {
		grade = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다.\n", score, grade)
}
