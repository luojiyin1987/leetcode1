func searchInsert(nums []int, target int) int {
    lenght  := len(nums)
    for i := 0 ; i< lenght ; i++ {
        if target  <= nums[i]  {
            return  i
        }
    }
return  lenght
}

//------------------------------
func searchInsert(nums []int, target int) int {

	for i, v := range nums {
		if target > v {
			continue
		}

		return i
	}

	return len(nums) 
}
