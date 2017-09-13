func maxSubArray(nums []int) int {
    lenght  := len(nums)
    if lenght == 1 {
        return nums[0]
    }
    b := nums[0]
    sum :=nums[0]
    for  i := 1; i<lenght ; i++ {
        if   nums[i] +b   < nums[i] {
               b =nums[i]
        }else {
            b = b+ nums[i]
        }
        if  sum < b {
            sum = b
        }
    }
    return sum
}
