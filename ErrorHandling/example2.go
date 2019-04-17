package main

import "fmt"

func main() {

	var answer1, answer2 string
	fmt.Print("Name:")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food:")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2)

}
