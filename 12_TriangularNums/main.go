package main

import (
	"fmt"
	"github.com/goldRust/ProjectEuler/dustutil"
)

func main() {
	var thisTri int
	for i := 1; ; i++ {

		thisTri += i

		factors := dustutil.GetFactors(thisTri)
		if len(factors) > 500 {
			fmt.Println(i, " : ", thisTri, " : ", factors)
			break
		}

	}
}
