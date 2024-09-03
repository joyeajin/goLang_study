package main

import "fmt"

func main() {
	age := 2

	switch age {
	case 1:
		fmt.Println("초등생1")
		fmt.Println("초등생2")
	case 2:
		fmt.Println("초등생3")
	default:
		fmt.Println("학생 아님")
	}

}
