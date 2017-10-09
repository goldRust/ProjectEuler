package main

import "fmt"

func main() {
	var sum int
	var sqsum int
	var sumsq int
	for i := 1; i <= 100; i++ {
		sumsq += i * i
	}

	for i := 1; i <= 100; i++ {
		sum += i
	}
	sqsum = sum * sum
	fmt.Println(sqsum - sumsq)
}
