package main

import "fmt"

func main() {
	qNum := 600851475143
	for i := 3; i < qNum; i++ {
		if isPrime(i) {
			if qNum%i == 0 {
				fmt.Println(i)
			}
		}

	}

}

func isPrime(num int) bool {
	for i := 3; i < num; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}
