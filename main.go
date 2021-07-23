package main

import (
	"fmt"
	"test_github_actions/cmd"
)

func main() {
	var x, y int
	x = 10
	y = 20

	fmt.Println("Program is starting.")

	sumResult := cmd.Sum(x, y)
	fmt.Printf("math result: %v\n", sumResult)

	mulResult := cmd.Mul(x, y)
	fmt.Printf("mul result: %v\n", mulResult)

	fmt.Println("Program ended with no errors.")

}
