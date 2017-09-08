package dustutil

import (
	"strconv"


)

func IsPalendrome(num int) bool{
	str:= strconv.Itoa(num)

	halfStr := len(str)/2
	for i:=0;i<halfStr;i++{
		if str[i]!=str[len(str)-(i+1)]{
			return false
		}
	}

	return true

}
func IsPrime(num int)bool{

	for i:=2;i<num/i;i++{
		if num%i==0{
			return false
		}

	}
	return true
}

func GetFactors(num int)[]int{
	factors := make([]int,0)

	for i:=1;i<num/i;i++{

		if num%i ==0{

			factors=append(factors,i)
			factors=append(factors,num/i)
		}
	}
	return factors
}

func InSliceInt(qSlice []int,x int)bool{
	for _,v:= range qSlice {
		if v==x{

			return true
		}

	}
	return false
}
func InSliceSlice(qSlice [][]int,x []int)bool{
	for _,v:= range qSlice {
		if len(v)!=len(x){continue}
		var count int
		for i:=0;i<len(v);i++{


			if v[i]!=x[i]{
				break
			}else{
				count++
				if count==len(v){return true}
			}
		}


		}






	return false
}