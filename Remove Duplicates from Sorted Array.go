func removeDuplicates(nums []int) int {
    if len(nums) <2 {
        return len(nums)
    }
    
    l := 0
    for i :=1; i < len(nums); i++ {
        if nums[i] != nums[l] {
            l++
            nums[l] = nums[i]
            
        }
    }
    return  l+1 
}
