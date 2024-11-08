package greeting

import "fmt"

func hi(name string) {
	fmt.Printf("Hi %s!\n", name)
}

func hello(name string) {
	fmt.Printf("Hello %s~\n", name)
}

func EnglishGreetings(name string) {
	hello(name)
	hi(name)
}
