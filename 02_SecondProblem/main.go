package main

import "fmt"

func main(){
 f1:=0
	f2:=1
	sum:=0


	for i:=0;i<4000000;{
		f1=f2
		f2 = i
		i=f1+f2

		if i%2==0{
			sum+=i
		}
	}
	fmt.Println(sum)
}
