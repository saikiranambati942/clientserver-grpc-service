package main

import "fmt"

// This go file tells us about different types in Go. we cannot run a GO program by having unused variables. It throws compilation error

func add(k, l float64) float64 {
	return k + l
}

func main() {
	var x int = 34
	var y float64 = float64(x)
	m := y // m will be type float64
	fmt.Println(add(m, y))

}
