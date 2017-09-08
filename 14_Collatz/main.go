package main

import "fmt"

func main(){
	var greatest  int
	var gChain []int
	for i:=1;i<1000000;i++{
		chain:= make([]int,0)
		thisI:=i
		for{
		if thisI==1{
				break
		}	else if thisI%2==0{
			thisI = thisI/2
			chain= append(chain,thisI)
		}else{
			thisI=thisI*3+1
			chain= append(chain,thisI)
			}


		}
	if len(chain)>len(gChain){
		gChain = chain
		greatest = i
	}
	}
	fmt.Println(greatest," : ", gChain)

}
