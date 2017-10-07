func checkPossibility(nums []int) bool {
    count :=0 
    for i :=1 ; i <len(nums);  i++  {
        if nums[i] < nums[i-1] {
            if count >= 1 {
                return false
            }
            count++ 
            if  i >1  && nums[i] <=nums[i-2]  {
                nums[i] = nums[i-1]
            }
        }
        
     
    } 
    return true
}
//----------------------------------------
func checkPossibility(nums []int) bool {
    var changes int
    
    for i := 0; i < len(nums)-1 && changes <= 1; i++ {
        if(nums[i] > nums[i+1]) {
            if(i < 1 || nums[i-1] <= nums[i+1]) {
                nums[i] = nums[i+1];
            } else {
                nums[i+1] = nums[i];
            }
            changes++;
        }
    }
    
    return changes <= 1;
}
