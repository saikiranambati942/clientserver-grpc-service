package main

import "fmt"

// This go file tells us about different types in Go. we cannot run a GO program by having unused variables. It throws compilation error

func add(x, y float64) float64 {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}

const x int = 5 // This is the way to define constants in GO
func main() {

	//var  num1 float64 = 3.5
	//var num2 float64 = 4.9
	//var num1,num2 float64=3.4,7.9  we can write in this format also
	//num1,num2 :=5.6,9.5 //we can also define variables in this way.Go will automatically find what type it is but type cannot be changed once it is compiled
	//fmt.Println(add(num1, num2))

	a, b := "hey", "sai"
	fmt.Println(multiple(a, b))

}
