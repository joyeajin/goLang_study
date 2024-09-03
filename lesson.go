package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Println("군대갈 수 있음")
	} else if age >= 14 {
		fmt.Println("학도병에 지원하세요")
	} else {
		fmt.Println("군대갈 수 없음")
	}

	// if age >= 18 {
	// 	fmt.Println("군대갈 수 있음")
	// }
	// if age >= 14 && age < 18 {
	// 	fmt.Println("학도병에 지원하세요")
	// }
	// if age < 14 {
	// 	fmt.Println("군대갈 수 없음")
	// }

}
