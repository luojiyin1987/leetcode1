package main

import "fmt"

func MaxSum(a []int) int{
	if len(a) == 1 {
		return a[0]
	}
	sum := a[0]
	b := a[0]
	for i :=1; i<len(a); i++{
		if a[i] > b+a[i] {
			b = a[i]
		}else{
			b = b+a[i]
		}
		if b > sum {
			sum = b
		}
	}
	fmt.Println(sum)
	return sum 
}

func main() {
	temp :=[]int{ 1,34,-20, -25,12,30,-23,45,12}
	MaxSum(temp)
}
