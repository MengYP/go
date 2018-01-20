package main

import "fmt"

func sum(a []int, result chan int)  {
	sum := 0
	for _;  v := range a {
		sum += v
	}
	result <- sum
}

func main()  {
	a := []int{2, 3, 5, 6, 10, -5, 1,}
}






