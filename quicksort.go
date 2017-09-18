package main

import (
	"fmt"
	"math/rand"
)

func qsort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivotIndex := rand.Int() % len(nums)

	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]

	qsort(nums[:left])
	qsort(nums[left+1:])

	return nums
}
func main() {
	nums := []int{345, 345, 45, 15, 45, 23, 45345, 23, 435, 556, 324324, 45546, 32434, 567657, 46567, 34234, 3465567, 34456, 45632, 456456}
	fmt.Println(qsort(nums))
}
