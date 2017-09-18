package main

import "fmt"

func Bubblesort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{4, 3, 20, 5, 90, 80, 10}
	fmt.Println(Bubblesort(nums))
}
