func reverseVowels(s string) string {
    //a,e,i,o,u 
     lenght := len(s)
    if lenght <2 {
        return s
    }
    var temp  []rune
    vowels := "aeiouAEIOU"
    vowelsMap := make(map[rune]int)
    for _,v  := range vowels {
        vowelsMap[v]++
    }
    for _,v := range s {
        _, ok := vowelsMap[v]
        if ok {
            temp = append(temp , v)
        }
    }
    
    turn :=1
    l := len(temp)
    fmt.Println(l)
    runes :=  []rune(s)
    for k, v := range runes {
        _, ok := vowelsMap[v]
        if ok {
           
            runes[k] = temp[l -turn]
            turn ++
            
         
        }
    }
    return string(runes)
    
}
 
