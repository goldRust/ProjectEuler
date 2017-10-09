package dustutil

import (
	"strconv"
)

//Any is an empty interface, allowing any TYPE to be identified as "Any"
type Any interface{}

//AllSlice is a slice of Anything - ints, strings, floats, etc...
type AllSlice []Any

//InSlice allows you to check a slice for any value. Returns bool.
//WARNING! Does not work on [][] - those can't be compared in this way.
func (as AllSlice) InSlice(qSlice []Any, x Any) bool {
	for _, v := range qSlice {
		if v == x {
			return true
		}
	}
	return false
}
func (as AllSlice) SliceSum() int {
	var sum int
	for _, v := range as {
		sum += (v).(int)
	}
	return sum
}

type SliceSlice []AllSlice

func (ss SliceSlice) InSlice(qSlice []AllSlice, x AllSlice) bool {
	for _, v := range qSlice {

		for i, w := range x {

			if w != v[i] {
				break
			}
			i++
			if i == len(v) {
				return true
			}
		}
	}
	return false
}

func IsPalendrome(num int) bool {
	str := strconv.Itoa(num)

	halfStr := len(str) / 2
	for i := 0; i < halfStr; i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}

	return true

}
func IsPrime(num int) bool {

	for i := 2; i < num/i; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}

func GetFactors(num int) []int {
	factors := make([]int, 0)

	for i := 1; i < num/i; i++ {

		if num%i == 0 {

			factors = append(factors, i)
			factors = append(factors, num/i)
		}
	}
	return factors
}


//InSliceInt searches a slice of ints for a particular int.
//Takes Slice of int, int. Index of int OR -1 if int not found.
func InSliceInt(qSlice []int,x int)int{
	for i,v:= range qSlice {
		if v==x{

			return i
		}

	}
	return -1

}

//InSliceSlice is not currently working.
func InSliceSlice(qSlice [][]int, x []int) bool {
	for _, v := range qSlice {
		if len(v) != len(x) {
			continue
		}
		var count int
		for i := 0; i < len(v); i++ {

			if v[i] != x[i] {
				break
			} else {
				count++
				if count == len(v) {
					return true
				}
			}
		}

	}

	return false
}


//InSliceString searches through a slice of strings for a particular string.
//Takes Slice of Stings, String. Returns index of x OR -1 if x is not found.
func InSliceString (qSlice []string, x string) int{
	for i,v:=range qSlice{
		if v == x{
			return i
		}
	}
	return -1
}
