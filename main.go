package main

import "fmt"

func main() {
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}
	// addFive := addN(5)
	result := addN(6)(1)
	//5 + 6 must print 7
	fmt.Println(result)
}
