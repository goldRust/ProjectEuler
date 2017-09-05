package dustutil

import "strconv"

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
	half:= num/2
	for i:=2;i<half;i++{
		if num%i==0{
			return false
		}

	}
	return true
}