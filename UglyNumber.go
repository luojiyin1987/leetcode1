func isUgly(num int) bool {
    for num > 1  {
        
        temp := num
        if  num%2 == 0 {
            num  = num /2
        }
        if num%3 == 0 {
            num = num  /3
        }
        if num % 5 == 0 {
            num = num /5
        }
        if num == temp {
            return false
        }
    }
    return num == 1
}
//-------------------------------------------------------
func isUgly(num int) bool {
    f := [3]int{2, 3, 5}
    if num > 0 {
        for _,n := range f {
            for num%n == 0 {
                num /= n
            }
        }   
    }
    return num == 1
}
