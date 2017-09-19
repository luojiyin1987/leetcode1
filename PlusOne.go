func plusOne(digits []int) []int {
    lenght := len(digits)
    result := make([]int, lenght+1, lenght+1)
    sum :=1 
    for i :=lenght -1 ; i >= 0 ; i-- {
        sum = sum + digits[i]
        result[i+1] = sum %10
        sum = sum /10
        if i==0 &&  sum >0  {
            result[0] = sum 
        }
    }
    if result[0] == 0 {
        return result[1:]
    }
    return result 
}
