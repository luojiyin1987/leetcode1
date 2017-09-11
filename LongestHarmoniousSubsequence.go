func findLHS(nums []int) int {
    numbers := make(map[int]int)
    for   _, v := range nums{
        numbers[v]++
    }
    var   temp []int
    for k,_:= range numbers {
        temp = append(temp, k)
    }
    sort.Ints(temp)
    length :=0
    for i := 0 ; i <len(temp) -1; i ++ {
        value := temp[i+1] - temp[i]
        if value == 1  &&  numbers[temp[i]] + numbers[temp[i+1]] > length {
            length  = numbers[temp[i]] + numbers[temp[i+1]]
        }
    }
    return  length    
}
//-----------------------------------------
func findLHS(nums []int) int {
    m := make(map[int]int, len(nums))
    for _, n := range nums {
        m[n]++
    }
    res := 0
    for n, count := range m {
        if v, ok := m[n+1]; ok && count+v > res {
            res = count+v
        }
    }
    return res
}
