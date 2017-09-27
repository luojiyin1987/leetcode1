func containsNearbyDuplicate(nums []int, k int) bool {
    if len(nums)  <2 {
        return false
    }
    nMap :=make(map[int]int)
    for i :=0; i<len(nums); i++ {
        _,ok := nMap[nums[i]] 
        if ok {
            idx := nMap[nums[i]] 
            if idx >=0 && i-idx <= k {
                return true 
            }
        }
        
        nMap[nums[i]] = i
    }
    return false
 }
 //------------------------------------------------------------------
 
