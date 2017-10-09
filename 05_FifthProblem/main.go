package main

import "fmt"

func main() {
	var num int
	i := 1
	//run loop until divBy returns true
	for {

		if divBy(i) {
			num = i
			break
		}

		i++
	}
	//Print the smallest multiple
	fmt.Println(num)

}
func divBy(num int) bool {
	//cycle through 1 - 20
	for i := 1; i <= 20; i++ {
		//eliminate any numbers that are not divisible by i
		if num%i != 0 {
			return false
		}
	}
	// return true if the int made it through
	return true
}
