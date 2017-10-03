//当  k== len(nums)  或者k=0 时，相当没翻转
//  当 k > len(nums)     k= k%ln  再处理


func rotate(nums []int, k int)  {
    ln := len(nums)
    if  ln <k {
        k =  k % ln
    }
    
    if ln > k && k >0 {
    
    ln_k := ln -k 
     
    for i, j := 0, ln_k-1; i< j;  i, j =i+1, j-1  {
        nums[i], nums[j] = nums[j], nums[i]
        //fmt.Println(1)
    }
    
    for i, j := ln_k, ln -1;  i<j; i,j = i+1 , j-1 {
        nums[i], nums[j] = nums[j], nums[i]
        //fmt.Println(2)
    }
    
    for i, j := 0, ln-1; i<j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
        //fmt.Println(3)
    }
    }
}
//----------------------------------------------
func rotate(nums []int, k int)  {
    size := len(nums)
    k = k % size
    if k <= 0 {
        return
    }
    for i := 0;i < size / 2; i += 1 {
        nums[i], nums[size - i - 1] = nums[size - i - 1], nums[i]
    }
    for i := 0;i < k / 2 ; i += 1 {
        nums[i], nums[k - i - 1] = nums[k - i - 1], nums[i]
    }
    for i := k; i < k + (size - k) / 2 ; i += 1 {
        nums[i], nums[size - i + k - 1] = nums[size - i + k - 1], nums[i]
    }
}
