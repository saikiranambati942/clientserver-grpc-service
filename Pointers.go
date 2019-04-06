package main

import "fmt"

func main() {
	a := 14
	x := &a //memory address of a
	fmt.Println("memory adress of a ", x)
	fmt.Println("value in the memory address of a ", *x)
	*x = 10
	fmt.Println(a)

	*x = *x * *x
	fmt.Println(*x)
	fmt.Println(a)
	fmt.Println(*&a)
}
