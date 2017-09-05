package main

import (
	"fmt"
	"github.com/goldRust/ProjectEuler/dustutil"
)


func main(){
	var sumPrime int =2
	for i:=3;i<=2000000;i++{
		if dustutil.IsPrime(i){
			sumPrime+=i
			fmt.Println("Added: ",i)

		}
	i++
	}
	fmt.Println(sumPrime)
}