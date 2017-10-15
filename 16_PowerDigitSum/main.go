package main

import (

	"fmt"

	"strconv"
)

func main() {
total:=1
	for i:=0;i<15;i++{
total*=2
	}
fmt.Println(total)
tots:=strconv.Itoa(total)
fmt.Println(tots)
var sum int
for i:=0;i<len(tots);i++{
	add,_:= strconv.Atoi(string(tots[i]))
	sum+=add
}
fmt.Println(sum)
}
