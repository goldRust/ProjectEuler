package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func main() {
	//Get the names file
	res, err := http.Get("https://projecteuler.net/project/resources/p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	//scan the file
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//prepare the math variables
	var bigMath int
	count := 0
	//scan the file
	for scanner.Scan() {
		word := scanner.Text()
		//split the names
		names := strings.Split(word, ",")
		//sort the names
		sort.Strings(names)
		for _, v := range names {
			count++
			var sum int
			sliceV := []byte(v)
			//remove the quotes.
			sliceV = append(sliceV[1 : len(sliceV)-1])
			v = string(sliceV)
			for _, l := range sliceV {
				//add up letter values
				sum += int(l) - 64
			}
			//math it out
			littleMath := sum * count
			bigMath += littleMath

		}

	}

	//Print final number
	fmt.Println(bigMath)
}
