func reverse(x int) int {
    temp := x 
    result := 0
    if x < 0{
        x = -x
    } 
     
    for x > 0  {
        result = result * 10 + x %10
        x = x /10
        //fmt.Println(result)
    }
    if   result  >= 2147483648 {
        return 0
    }
    if  temp <0 {
    result = -result 
    }
    return  result
    
}
