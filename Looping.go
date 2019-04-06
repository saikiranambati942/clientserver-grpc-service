package main

import "fmt"

func main() {
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	/*i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}*/

	/*for{
		fmt.Println("Hey Sai")     This is like an infinite while loop
	}*/
	i := 5

	for {
		fmt.Println("Hey Sai", i)
		i += 3
		if i > 25 {
			break
		}
	}

	/*for i := 5; i < 25; i += 3 {
		fmt.Println("Hey Sai", i)   we can write in this way also
	}*/
}
