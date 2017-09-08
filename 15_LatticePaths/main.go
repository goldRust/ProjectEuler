package main

import (
	"fmt"

)


func main(){
	grid:= make([][]int,16)

	path:= 0
	for i:=0;i<2;i++{
		for j:=0;j<2 ;j++  {
			grid[path]=append(grid[path],j)
		}
	}
	fmt.Println(grid)
}
