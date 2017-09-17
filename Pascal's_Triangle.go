func generate(numRows int) [][]int {
    //seed := 1
    var temp1 [][]int
     
    for i :=1 ; i <=numRows ;  i++ {
        var  temp2  []int
        for j :=0 ; j < i ; j++ {
            if j== 0  || j == i -1 {
                temp2 = append(temp2, 1)
            } else   {
                t  := temp1[i-2][j-1] + temp1[i-2][j]
                temp2  = append(temp2, t)
            }
            fmt.Println(temp2)
        }
        temp1 = append(temp1, temp2)
    }
    return  temp1
}
