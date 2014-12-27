package main

import "fmt"

func f(from string, l int) {
	for i := 0; i < l; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("goroutine", 3)

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	f("direct", 3)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
