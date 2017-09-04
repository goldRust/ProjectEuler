package main

import "fmt"

func main(){
	var num int
	i:=1
	for {

		if divBy(i){num=i
		break}

		i++
	}
	fmt.Println(num)

}
	func divBy(num int)bool{
		for i:=1;i<=20;i++{
			if num%i!=0{
				return false
			}
		}
		return true
	}
