package main

import "fmt"

func main() {

	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
}

//func (r receiver) identifier(parameters) (return (s)){...}

func foo() {
	fmt.Println("hello from foo")
}

//Everything in GO is PASS BY VALUE
func bar(s string) {
	fmt.Println(s)
}

func woo(st string) string {
	return fmt.Sprint("Hello ", st)
}
