package main

import "fmt"

func MaxSum(a []int) int{
	sum := 0
	b := 0
	for i :=0; i<len(a); i++{
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
