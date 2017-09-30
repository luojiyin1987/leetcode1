func thirdMax(nums []int) int {
    first, second , third := -2147483648, -2147483648, -2147483648
    
    temp  := make(map[int]int) 
    
    for _, v := range nums {
        temp[v]++
    }
    
    for  i :=0; i<len(nums); i++ {
        if third >= nums[i] || first == nums[i] || second== nums[i] {
            continue
        }
        third = nums[i] 
        if second  < third {
            second , third = third, second
        }
        if first < second {
            first , second = second , first
        }
    }
    if len(temp) <=2 {
        return first
    }
    return third 
    
}
//不能像Pytohon 那样那样把值初始化为  None，不断比较，更新 first second   third 的值， 最后，如果third 为None，
没有第三大值，返回first，third不为 None， 则返回third
所以要把值初始化为-2147483648  ， int32中 最小值，然后 
