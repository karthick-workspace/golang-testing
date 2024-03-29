package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//To test commit
	add := adder()
	fmt.Println(add(10))
}
