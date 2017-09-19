func isPerfectSquare(num int) bool {
    sum :=0 
    turn :=0
    for  sum  <num {
        sum =  2 * turn + 1 +sum
        turn ++ 
      
    }
    return  sum == num 
}
