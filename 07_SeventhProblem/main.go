package main

import "fmt"

func main() {
	counter := 0
	i := 1
	for {
		i++
		if isPrime(i) {
			counter++
			fmt.Println(counter, " - ", i)
		}
		if counter == 10001 {
			fmt.Print(i)
			break
		}
	}
}
func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}
