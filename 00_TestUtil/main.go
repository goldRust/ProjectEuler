package main

import (
	"fmt"
	"github.com/goldRust/ProjectEuler/dustutil"
)

func main() {
	var si = dustutil.AllSlice{3, 3, 4, 2, 4, 3, 1, 9, 4, 34, 64, 11, 3, 433, 545, 43, "i"}
	for i := 0; i < 700; i++ {
		if si.InSlice(si, i) {
			fmt.Println(i, " is in the slice!")
		}
	}

	var ss = dustutil.AllSlice{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

	for i := 55; i < 200; i++ {
		letter := string(i)
		if ss.InSlice(ss, letter) {
			fmt.Println(letter, "is in the slice !")
		}
	}
	var superslice = dustutil.SliceSlice{ss, si}
	var notThere = dustutil.AllSlice{3, 3, 4, 2, 4}
	if superslice.InSlice(superslice, notThere) {
		fmt.Println("Found the slice in the slice!")
	} else {
		fmt.Println("Slice not found.")
	}
	superslice = append(superslice, notThere)
	if superslice.InSlice(superslice, notThere) {
		fmt.Println("Oh, there it is.")
	} else {
		fmt.Println("Nope, still not there.")
	}

}
