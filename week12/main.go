package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	fmt.Println(scores[1], scores[0]) //, scores[3])  // invalid argument: index 3 out of bounds
	fmt.Printf("%#v\n", scores)
	fmt.Println(scores)

	// var dates [3]time.Time
	// dates[1] = time.Unix(1947200001, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1947200001, 0), // need comma
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947200001, 0)}
	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Printf("%#v\n", dates)
	fmt.Println(dates)
}
