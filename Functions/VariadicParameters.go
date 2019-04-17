package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6)	

	foo(1, 3, 4)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}