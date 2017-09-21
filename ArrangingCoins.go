func arrangeCoins(n int) int {
    turn := 0
    for i :=1 ; i<= n ; i++ {
        n = n -i 
        if n >= 0 {
            turn ++
        }
    }
    return turn 
}
