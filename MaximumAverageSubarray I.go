func findMaxAverage(nums []int, k int) float64 {
    
    sum := 0
    for i :=0; i<k ; i++ {
        sum = sum + nums[i]
      
    }
  
    max := sum
    for  j := k; j<len(nums) ; j++ {
        sum = sum + nums[j] - nums[j-k]
    
        if sum > max {
            max = sum
        }
    }
    return    float64(max)/(float64(k))
}
