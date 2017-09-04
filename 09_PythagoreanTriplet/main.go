package main

import "fmt"

func main(){
	these := make([]int,3)
	var prod int
 for i:=0;i<1000;i++{
	 for j:=0;j<i;j++{
		 for k:=0;k<j;k++{
			 if k+j+i==1000 {
				 i2 := i * i
				 j2 := j * j
				 k2 := k * k
				 if k2+j2==i2{
					 prod = i*j*k
					 these[0]=k
					 these[1]=j
					 these[2]=i
					 break

				 }
			 }
		 }
	 }
 }
	fmt.Println("triplet found! - ",these," -- Product: ",prod)
}
