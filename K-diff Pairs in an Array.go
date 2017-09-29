func findPairs(nums []int, k int) int {
    if k  <0 {
        return 0
    }
    temp := make(map[int]int)
    ret := 0
    for _,v :=range nums {
        temp[v]++
    }
    //fmt.Println(temp)
    
    for num, v :=range temp {
        _, ok := temp[num+k] 
        if k >0 {
            if ok {
                ret++
            }
        }
        if k == 0 {
            if ok && v >=2{
                ret++
            }
        }
    }
    return ret 
}
//-------------------------------------

func findPairs(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }
    
    c := 0
    for num, cnt := range m {
        if k == 0 {
            if cnt > 1 {
                c++
            }
        } else {
            diff := num - k
            if m[diff] > 0 {
                    c++
                }
        }
    }
    
    return c
}
