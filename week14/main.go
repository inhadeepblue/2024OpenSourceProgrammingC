package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	// visitors : 구조체 슬라이스
	totalCost := 0
	for _, v := range visitors {
		totalCost = totalCost + v.cost
	}
	return totalCost
}

func main() {
	var numVisitors int
	fmt.Println("How many visitors? ")
	fmt.Scanln(&numVisitors)

	vs := make([]visitor, numVisitors) // create Slice

	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Print("Input age : ")
		fmt.Scan(&age)

		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 5000}
		default:
			vs[i] = visitor{age: age, cost: 7000}
		}
	}
	fmt.Printf("Total price is %d won!", calculateCost(vs))
}
