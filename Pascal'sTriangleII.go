func getRow(rowIndex int) []int {
    ret := []int{1}
    if rowIndex == 0{
        return ret
    }
    for i:=1;i<=rowIndex;i++{
        ret = append([]int{1},ret...)
        for j:=1;j<i;j++{
            ret[j] = ret[j] + ret[j+1]
        }
    }
    return ret
}

