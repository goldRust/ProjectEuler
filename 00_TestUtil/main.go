package main

import (
	"github.com/goldRust/ProjectEuler/dustutil"
	"fmt"
)

func main() {
	var si= dustutil.IntSlice{3,3,4,2,4,3,1,9,4,34,64,11,3,433,545,43,}
	for i:=0;i<700;i++{
		if si.InSlice(si,i){
		fmt.Println(i," is in the slice!")}
	}

	var ss = []string{"a","b","c","d","e","f","g","h","i","j","k","l"}

	for i:=55;i<200;i++{
		letter:= string(i)
		if dustutil.StringSlice(ss).InSlice(ss,letter) {
			fmt.Println(letter, "is in the slice !")
		}
	}

}
