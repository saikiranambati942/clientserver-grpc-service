package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("OS\t",runtime.GOOS)
	fmt.Println("ARCH\t",runtime.GOARCH)
	fmt.Println("CPU's\t",runtime.NumCPU())
	fmt.Println("Goroutines\t",runtime.NumGoroutine())
	foo()
	bar()
}

func foo(){
	for i:=0;i<10;i++{
		fmt.Println("foo: ",i)
	}
}

func bar(){
	for i:=0;i<10;i++{
		fmt.Println("bar: ",i)
	}
}