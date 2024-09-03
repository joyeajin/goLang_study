package main

import "fmt"

func main() {
	// x := 1
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }
	for x := 1; x <= 15; x++ {
		// fmt.Println(x)
		if x == 3 {
			continue
		} else if x == 8 {
			break
		}
		fmt.Println(x)
	}
}
