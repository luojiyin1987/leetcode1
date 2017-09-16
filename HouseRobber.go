func rob(nums []int) int {
    lenght  := len(nums)
    if lenght == 1  {
        return nums[0]
    }
    
    last , now := 0, 0
    
    for _, i  := range nums {
        t := now
        last += i 
        if last > now {
            now = last
        }
        last = t 
    }
    return now 
}
