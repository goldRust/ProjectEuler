package main

import (
	"fmt"
)

//Starting over with Pascal's triangle


func main(){
	//n is equal to total number of steps (length + width of grid)
	n:=40
	//This tracks the previous level of pascal's triangle - needed to calculate the next level.
	var lastLine []int
	for i:=0;i<n;i++{
		var pascSlice []int
		pascSlice=append(pascSlice,1)
		for j:=0;j<len(lastLine)-1;j++{
			calc:=lastLine[j]+lastLine[j+1]
			pascSlice= append(pascSlice,calc)
		}
		pascSlice=append(pascSlice,1)
		lastLine=pascSlice
		fmt.Println(lastLine)
	}
 mid:=lastLine[len(lastLine)/2]
 fmt.Println(mid)
}
//BOOM! IT WORKS!