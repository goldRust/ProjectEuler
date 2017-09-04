package main

import (
	"fmt"
	"strconv"
)
var great int
func main(){
	var pals int
	for i:=100;i<1000;i++{
		for j:=100;j<1000;j++{
			if isPalendrome(i*j){
				pals = isGreatest(i*j)
			}
		}
	}
 fmt.Println(pals)
}
func isPalendrome(num int) bool{
    str:= strconv.Itoa(num)

	halfStr := len(str)/2
	for i:=0;i<halfStr;i++{
		if str[i]!=str[len(str)-(i+1)]{
			return false
		}
	}

	return true

}
func isGreatest(num int)int {

		if num>great{
			great=num
		}

	return great
}
