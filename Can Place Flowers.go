func canPlaceFlowers(flowerbed []int, n int) bool {
    count :=0
    for i := 0 ; i <len(flowerbed); i++ {
        if flowerbed[i] ==1 {
            continue
        }
        if i>0 && flowerbed[i-1] ==1 {
            continue
        } 
        if i <len(flowerbed) -1 && flowerbed[i+1] ==1 {
            continue
        }
        count ++
        flowerbed[i] =1
    }
    return count >= n
    
}
